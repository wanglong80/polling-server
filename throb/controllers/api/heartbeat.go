package api

import (
	"encoding/json"
	"happy.work/throb/service"
)

type HeartbeatController struct {
	BaseController
}

// 轮询心跳包，代替长连接，时间间隔短用来获取一些需要实时变化的数据
// 这里提供的是内存级读取，为了减少 DB 操作，提高 QPS
func (c *HeartbeatController) Index() {
	req := make(map[string]string)
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	ms := make(map[string]interface{})

	// 获取消息数据
	for index, id := range req {
		// 代表仅取最新的一条消息
		if id == "0" {
			ms[index] = service.MessageLastCache(index)
		} else {
			ms[index] = service.MessageListCache(index, id)
		}
	}

	c.Data["json"] = ResponseWrapper{Code: 0, Message: "OK", Data: ms}
	c.ServeJSON()
}
