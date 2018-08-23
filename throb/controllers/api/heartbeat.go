package api

import (
	"happy.work/throb/requests"
	"happy.work/throb/service"
	"strconv"
)

type HeartbeatController struct {
	BaseController
}

// 轮询心跳包，代替长连接，时间间隔短用来获取一些需要实时变化的数据
// 这里提供的是内存级读取，为了减少 DB 操作，提高 QPS
func (c *HeartbeatController) Index() {
	req := &requests.HeartbeatReq{}
	message, err := c.RequestData(req)

	if err != nil {
		c.Response(0, message, nil)
		return
	}

	ms := make(map[string]interface{})
	ss := make(map[string]interface{})

	// 获取消息数据
	for _, e := range req.Ms {
		// gtId参数等于 -1 代表仅取最新的一条消息
		if e.Last {
			ms[e.Name] = service.MessageLastCache(e.Name)
		} else {
			var gtIdStr string
			var ltIdStr string

			if e.GtId == 0 {
				gtIdStr = "-inf"
			} else {
				gtIdStr = strconv.FormatInt(e.GtId, 10)
			}

			if e.LtId == 0 {
				ltIdStr = "+inf"
			} else {
				ltIdStr = strconv.FormatInt(e.LtId, 10)
			}

			ms[e.Name] = service.MessageListCache(e.Name, gtIdStr, ltIdStr)
		}
	}

	// 获取状态数据
	for _, name := range req.Ss {
		ss[name] = service.StorageCache(name)
	}

	resp := make(map[string]interface{})

	// 消息空间，存储各消息空间的消息序列（有序列表并可选持久化存储）
	resp["ms"] = ms
	// 状态空间，数据与状态存储器
	resp["ss"] = ss

	c.Data["json"] = ResponseWrapper{Code: 0, Message: "OK", Data: resp}
	c.ServeJSON()

}
