package api

import (
	"happy.work/throb/requests"
	"happy.work/throb/service"
)

type StorageController struct {
	BaseController
}

// 获取状态数据
func (c *StorageController) GetStorage() {
	name := c.GetString("name")
	m := service.StorageCache(name)
	c.Data["json"] = ResponseWrapper{Code: 0, Message: "OK", Data: m}
	c.ServeJSON()
}

// 创建状态数据
func (c *StorageController) CreateStorage() {
	req := &requests.CreateStorageReq{}
	message, err := c.RequestData(req)

	if err != nil {
		c.Response(0, message, nil)
		return
	}

	m := service.CreateStorage(req.Key, req.Data)
	c.Data["json"] = ResponseWrapper{Code: 0, Message: "OK", Data: m}
	c.ServeJSON()
}