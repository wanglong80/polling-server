package api

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// 统一响应返回体
type ResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 获取并校验请求数据
func (c *BaseController) RequestData(req interface{}) (string, error) {
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	//valid := validation.Validation{}
	//ok, _ := valid.Valid(req)

	//if !ok {
	//	beego.Debug(valid)
	//	message := "err"
	//	c.Data["json"] = ResponseWrapper{Code: 1, Message: message, Data: ok}
	//	c.ServeJSON()
	//
	//	return message, errors.New("InvalidParameter")
	//}

	return "ok", nil
}

// 输出响应数据
func (c *BaseController) Response(code int, message string, data interface{}) {
	c.Data["json"] = ResponseWrapper{Code: code, Message: message, Data: data}
	c.ServeJSON()
}
