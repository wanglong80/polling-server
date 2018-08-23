package api

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "RlfHIi75FKei2RFsNjwrXgJ8Mj3O6iba"
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
func  (c *BaseController) RequestData(req interface{}) (string, error) {
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	valid := validation.Validation{}
	ok, _ := valid.Valid(req)

	if !ok {
		beego.Debug(valid)
		message := valid.Errors[0].Key + " " +  valid.Errors[0].Message
		c.Data["json"] = ResponseWrapper{Code: 1, Message: message, Data: ok}
		c.ServeJSON()

		return message, errors.New("InvalidParameter")
	}

	return "ok", nil
}

// 获取 JWT Token
func  (c *BaseController) JwtToken() (jwt.MapClaims, error) {
	token, err := request.ParseFromRequest(c.Ctx.Request, request.AuthorizationHeaderExtractor,
	func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("InvalidToken")
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}

// 输出响应数据
func  (c *BaseController) Response(code int, message string, data interface{}) {
	c.Data["json"] = ResponseWrapper{Code: code, Message: message, Data: data}
	c.ServeJSON()
}