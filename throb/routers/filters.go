package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/astaxie/beego"
)

// 登录校验
func auth(ctx *context.Context) {
	token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if ctx.Request.Method == "OPTIONS" {
		return
	}

	claims := token.Claims.(jwt.MapClaims)


	beego.Debug(claims["sub"])

	if err != nil || !token.Valid {
		ctx.Output.Header("Access-Control-Allow-Headers", "Origin,Authorization,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		ctx.Output.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		ctx.Abort(401, "401")
	}
}
