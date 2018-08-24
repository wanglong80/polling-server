package api

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["json"] = ResponseWrapper{Code: 404, Message: "请求的资源不存在"}
	c.ServeJSON()
}

func (c *ErrorController) ErrorUnauthorized() {
	c.Data["json"] = ResponseWrapper{Code: 1001, Message: "身份认证失败"}
	c.ServeJSON()
}

func (c *ErrorController) ErrorTokenInvalid() {
	c.Data["json"] = ResponseWrapper{Code: 1002, Message: "无效的授权令牌或已过期"}
	c.ServeJSON()
}

func (c *ErrorController) ErrorHeaderMissing() {
	c.Data["json"] = ResponseWrapper{Code: 1003, Message: "缺少必须的头信息"}
	c.ServeJSON()
}