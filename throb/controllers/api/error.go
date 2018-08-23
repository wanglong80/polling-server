package api

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["json"] = ResponseWrapper{Code: 404, Message: "请求的资源不存在"}
	c.ServeJSON()
}

func (c *ErrorController) Error401() {
	c.Data["json"] = ResponseWrapper{Code: 401, Message: "身份认证失败"}
	c.ServeJSON()
}

func (c *ErrorController) Error412() {
	c.Data["json"] = ResponseWrapper{Code: 412, Message: "缺少必须的头信息"}
	c.ServeJSON()
}
