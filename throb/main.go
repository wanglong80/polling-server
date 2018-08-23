package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/fatih/color"
	"happy.work/throb/controllers/api"
	_ "happy.work/throb/routers"
)

var UserId string

func main() {
	slogan()
	config()
	errors()
	beego.Run()
}

// 初始化配置文件
func config() {
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.AppName = "throb"
	beego.BConfig.ServerName = "throb"
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.RunMode = beego.AppConfig.String("app.env")
	beego.BConfig.EnableGzip, _ = beego.AppConfig.Bool("app.enable_gzip")
	beego.BConfig.Listen.HTTPPort, _ = beego.AppConfig.Int("app.http_port")
	beego.BConfig.Listen.EnableAdmin = false

}

// 设置错误控制器
func errors() {
	beego.ErrorController(&api.ErrorController{})
}

// 欢迎字符图形

func slogan() {
	/*
	 * http://www.network-science.de/ascii
	 * Font:standard
	 */
	ascii := `
 _____ _               _
|_   _| |__  _ __ ___ | |__
  | | | '_ \| '__/ _ \| '_ \
  | | | | | | | | (_) | |_) |
  |_| |_| |_|_|  \___/|_.__/  v1.0
`

	color.Set(color.FgHiMagenta, color.Bold)
	defer color.Unset()
	fmt.Println(ascii)
}
