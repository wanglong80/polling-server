package routers

import (
	"github.com/astaxie/beego"
	"happy.work/throb/controllers/api"
)

func init() {
	ns :=
		beego.NewNamespace("/api",
			// 心跳包
			NSRoute("/heartbeat", &api.HeartbeatController{}, "post:Index", auth),

			// 消息
			NSRoute("/message/list", &api.MessageController{}, "post:GetMessageList"),
			NSRoute("/message/create", &api.MessageController{}, "post:CreateMessage", auth),
			NSRoute("/message/delete", &api.MessageController{}, "post:DeleteMessage"),
			NSRoute("/index/delete", &api.MessageController{}, "post:DeleteIndex"),

			// 存储
			NSRoute("/storage/get", &api.StorageController{}, "get:GetStorage"),
			NSRoute("/storage/create", &api.StorageController{}, "post:CreateStorage"),
		)

	beego.AddNamespace(ns)
}
