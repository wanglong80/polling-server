package service

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"happy.work/throb/cache"
	"happy.work/throb/db"
	"happy.work/throb/requests"
	"time"
)

type SimpleMessageItem struct {
	Id        int64  `json:"i"`
	Type      int64  `json:"t"`
	Uid       string `json:"u"`
	Body      string `json:"b"`
	CreatedAt int64  `json:"c"`
}

// 获取最大ID的消息（直接读缓存，提高QPS）
func MessageLastCache(name string) []SimpleMessageItem {
	r := cache.Conn()

	if r == nil {
		return nil
	}

	defer r.Close()

	key := "imx__" + name

	values := cache.GetListValuesByMaxId(r, key)

	if values == nil {
		return nil
	}

	var simpleMessageItems []SimpleMessageItem

	for _, v := range values {
		var m SimpleMessageItem
		json.Unmarshal([]byte(v.([]byte)), &m)
		simpleMessageItems = append(simpleMessageItems, m)
	}

	return simpleMessageItems
}

// 获取消息列表（直接读缓存，提高QPS）
func MessageListCache(name string, gtIdStr string, ltIdStr string) []SimpleMessageItem {
	r := cache.Conn()

	if r == nil {
		return nil
	}

	defer r.Close()

	key := "imx__" + name
	minId := "(" + gtIdStr
	maxId := "(" + ltIdStr

	values := cache.GetListValuesByRange(r, key, minId, maxId)

	if values == nil {
		return nil
	}

	var simpleMessageItems []SimpleMessageItem

	for _, v := range values {
		var m SimpleMessageItem
		json.Unmarshal([]byte(v.([]byte)), &m)
		simpleMessageItems = append(simpleMessageItems, m)
	}

	return simpleMessageItems
}

// 获取历史消息列表
func GetMessageList(filter *requests.GetMessageListReq) []*SimpleMessageItem {
	o := orm.NewOrm()
	qs := o.QueryTable("message")

	qs = qs.Filter("index", filter.Index)

	if filter.Limit > 0 {
		qs = qs.Limit(filter.Limit)
	}

	if filter.Offset > 0 {
		qs = qs.Limit(filter.Offset)
	}

	if filter.Offset > 0 && filter.Limit > 0 {
		qs = qs.Limit(filter.Limit, filter.Offset)
	}

	qs = qs.OrderBy("-id")

	var messages []*db.Message
	qs.All(&messages)

	var simpleMessageItems []*SimpleMessageItem

	for _, v := range messages {
		var m SimpleMessageItem
		m.Id = v.Id
		m.Uid = v.Uid
		m.Type = v.Type
		m.Body = v.Body
		m.CreatedAt = v.CreatedAt
		simpleMessageItems = append(simpleMessageItems, &m)
	}

	return simpleMessageItems
}

// 创建消息
func CreateMessage(message *db.Message, isPersistence bool) int64 {
	message.CreatedAt = time.Now().Unix()

	var id int64

	// 消息持久化，可以查询历史记录
	if isPersistence {
		o := orm.NewOrm()
		insertId, err := o.Insert(message)

		if err != nil {
			logs.Error(err)
		}
		id = insertId
	} else {
		var originTimestamp int64 = 1429164800
		// 临时内存的方式因为无法得知递增唯一ID，只能用时间戳来表示（是否唯一不重要）
		// 之所以减掉一个基础时间戳只是为了让这个数的长度尽可能的变短，减少内存使用
		id = time.Now().Unix() - originTimestamp
	}

	messageListItem := &SimpleMessageItem{}
	messageListItem.Id = id
	messageListItem.Type = message.Type
	messageListItem.Uid = message.Uid
	messageListItem.Body = message.Body
	messageListItem.CreatedAt = message.CreatedAt

	// 把数据放入内存
	r := cache.Conn()

	if r == nil {
		return 0
	}

	defer r.Close()

	key := "imx__" + message.Index

	if id > 0 {
		str, _ := json.Marshal(messageListItem)
		cache.AddListValue(r, key, id, string(str))
	}

	return id
}

// 删除消息
func DeleteMessage(index string, id int64) int64 {
	// 数据库中删除数据
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM message WHERE `index` = ? AND `id` = ?", index, id).Exec()

	if err == nil {
		num, _ := res.RowsAffected()

		key := "imx__" + index

		if num > 0 {
			// 内存中删除数据
			r := cache.Conn()

			if r == nil {
				return 0
			}

			defer r.Close()

			cache.DeleteListValueById(r, key, id)
		}
		return num
	}

	return 0
}

// 删除消息索引
func DeleteIndex(index string) int64 {
	// 数据库中删除数据
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM message WHERE `index` = ?", index).Exec()

	if err == nil {
		num, _ := res.RowsAffected()

		key := "imx__" + index

		if num > 0 {
			// 内存中删除数据
			r := cache.Conn()

			if r == nil {
				return 0
			}

			defer r.Close()

			cache.DeleteKey(r, key)
		}
		return num
	}

	return 0
}
