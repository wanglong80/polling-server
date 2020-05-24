package service

import (
	"encoding/json"
	"happy.work/throb/cache"
	"happy.work/throb/requests"
	"strconv"
	"time"
)

type SimpleMessageItem struct {
	Id        string  `json:"i"`
	Uid       string `json:"u"`
	Body      string `json:"b"`
	CreatedAt int64  `json:"c"`
}

// 获取最大ID的消息（直接读缓存，提高QPS）
func MessageLastCache(index string) []SimpleMessageItem {
	r := cache.Conn()

	if r == nil {
		return nil
	}

	defer r.Close()

	key := "imx__" + index

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
func MessageListCache(index string, id string) []SimpleMessageItem {
	r := cache.Conn()

	if r == nil {
		return nil
	}

	defer r.Close()

	key := "imx__" + index
	minId := "(" + id
	maxId := "(+inf"

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

// 创建消息
func CreateMessage(message *requests.CreateMessageReq) string {
	var originTimestamp int64 = 1590309350884000000
	// 临时内存的方式因为无法得知递增唯一ID，只能用时间戳来表示
	// 之所以减掉一个基础时间戳只是为了让这个数的长度尽可能的变短，减少内存使用
	// 一旦部署不建议修改此值，除非清空全部缓存
	var id = time.Now().UnixNano() - originTimestamp
	var idStr = strconv.FormatInt(id, 10)

	messageListItem := &SimpleMessageItem{}
	messageListItem.Id = idStr
	messageListItem.Uid = message.Uid
	messageListItem.Body = message.Body
	messageListItem.CreatedAt = time.Now().Unix()

	// 把数据放入内存
	r := cache.Conn()

	if r == nil {
		return "0"
	}

	defer r.Close()

	key := "imx__" + message.Index
	str, _ := json.Marshal(messageListItem)
	cache.AddListValue(r, key, id, string(str))

	return idStr
}

// 删除消息
func DeleteMessage(index string, id int64) int64 {
	key := "imx__" + index

	r := cache.Conn()

	if r == nil {
		return 0
	}

	defer r.Close()

	cache.DeleteListValueById(r, key, id)

	return 0
}

// 删除消息索引
func DeleteIndex(index string) int64 {
	key := "imx__" + index

	r := cache.Conn()

	if r == nil {
		return 0
	}

	defer r.Close()

	cache.DeleteKey(r, key)

	return 0
}
