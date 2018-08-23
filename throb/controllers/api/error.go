package api

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["json"] = ResponseWrapper{Code: 404, Message: "请求的资源不存在"}
	c.ServeJSON()
}


func (c *ErrorController) Error401() {
	c.Data["json"] = ResponseWrapper{Code: 401, Message: "身份认证过期或未登录"}
	c.ServeJSON()
}