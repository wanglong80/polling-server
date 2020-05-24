package api

import (
	"happy.work/throb/global"
	"happy.work/throb/requests"
	"happy.work/throb/service"
)

type MessageController struct {
	BaseController
}

// 创建消息
func (c *MessageController) CreateMessage() {
	req := &requests.CreateMessageReq{}
	msg, err := c.RequestData(req)

	if err != nil {
		c.Response(0, msg, nil)
		return
	}

	message := &requests.CreateMessageReq{}
	message.Index = req.Index
	message.Body = req.Body
	message.Uid = global.UserId

	id := service.CreateMessage(message)

	c.Data["json"] = ResponseWrapper{Code: 0, Message: "OK", Data: id}
	c.ServeJSON()
}

// 删除消息
func (c *MessageController) DeleteMessage() {
	req := &requests.DeleteMessageReq{}
	message, err := c.RequestData(req)

	if err != nil {
		c.Response(0, message, nil)
		return
	}

	num := service.DeleteMessage(req.Index, req.Id)

	c.Data["json"] = ResponseWrapper{Code: 0, Message: "OK", Data: num}
	c.ServeJSON()
}

// 删除索引
func (c *MessageController) DeleteIndex() {
	req := &requests.DeleteIndexReq{}
	message, err := c.RequestData(req)

	if err != nil {
		c.Response(0, message, nil)
		return
	}

	num := service.DeleteIndex(req.Index)

	c.Data["json"] = ResponseWrapper{Code: 0, Message: "OK", Data: num}
	c.ServeJSON()
}
