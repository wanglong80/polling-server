package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"happy.work/throb/global"
	"time"
)

const (
	JWTSecret = "RlfHIi75FKei2RFsNjwrXgJ8Mj3O6iba"
	MasterKey = "NsyzoP99Dsk3RXUUtqxjwo8Mj312lqwr"
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
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "User-Id"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		MaxAge: time.Hour,
	}))

	// 鉴权验证
	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		userId := ctx.Request.Header.Get("User-Id")
		masterKey := ctx.Request.Header.Get("Master-Key")

		token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(JWTSecret), nil
			})

		if ctx.Request.Method == "OPTIONS" {
			return
		}

		// User-Id 必传
		if len(userId) == 0 {
			ctx.Abort(412, "HeaderMissing")
		}

		// Master-Key 不存在并且开启了 JWT 校验则需要判断 token
		jwtEnabled, _ := beego.AppConfig.Bool("jwt.enabled")

		if len(masterKey) == 0 || masterKey != MasterKey {
			if jwtEnabled {
				if token == nil {
					ctx.Abort(401, "Unauthorized")
				}

				claims := token.Claims.(jwt.MapClaims)

				// JWT Token 中缺少 sub
				if claims["sub"] == nil {
					ctx.Abort(401, "Unauthorized")
				}

				// sub 中的值和 userId 不一致，鉴权不通过
				sub := fmt.Sprintf("%v", claims["sub"])
				if userId != sub {
					ctx.Abort(401, "Unauthorized")
				}
			}
		}

		if err != nil || !token.Valid {
			ctx.Abort(401, "TokenInvalid")
		}

		global.UserId = userId
	})

	// 每次请求重置全局变量
	beego.InsertFilter("*", beego.BeforeStatic, func(ctx *context.Context) {
		global.UserId = ""
	})
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
