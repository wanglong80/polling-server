package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

const (
	SecretKey = "RlfHIi75FKei2RFsNjwrXgJ8Mj3O6iba"
)

func init() {
	// 不是 API 的路由统一走向一个默认路由，该路由指向一个纯前端项目
	// 前端是 SPA 应用，路由控制权交给前端，因此所有前缀不是 API 的 URL 都应该重写至首页
	//beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
	//	match, _ := regexp.MatchString("^/api/", ctx.Input.URI())
	//
	//	ctx.Input.RunController = nil
	//	ctx.Input.RunMethod = ""
	//
	//	if !match {
	//		ctx.Input.RunController = reflect.TypeOf(controllers.MainController{})
	//		ctx.Input.RunMethod = "Index"
	//	}
	//})

	// 跨域设置
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}


func NSRoute(rootpath string, c beego.ControllerInterface, method string, filterList ...beego.FilterFunc) beego.LinkNamespace {
	return func(ns *beego.Namespace) {
		n := beego.NewNamespace(rootpath,
			beego.NSBefore(filterList...),
			beego.NSRouter("/", c, method),
		)
		ns.Namespace(n)
	}
}