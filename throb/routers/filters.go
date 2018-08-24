package routers



// 登录校验
//func auth(ctx *context.Context) {
//	userId := ctx.Request.Header.Get("User-Id")
//	masterKey := ctx.Request.Header.Get("Master-Key")
//
//	token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
//		func(token *jwt.Token) (interface{}, error) {
//			return []byte(JWTSecret), nil
//		})
//
//	if ctx.Request.Method == "OPTIONS" {
//		return
//	}
//
//	// User-Id 必传
//	if len(userId) == 0 {
//		ctx.Abort(412, "412")
//	}
//
//	// Master-Key 不存在并且开启了 JWT 校验则需要判断 token
//	jwtEnabled, _ := beego.AppConfig.Bool("jwt.enabled")
//
//	if len(masterKey) == 0 || masterKey != MasterKey {
//		if jwtEnabled {
//			if token == nil {
//				ctx.Abort(401, "401")
//			}
//
//			claims := token.Claims.(jwt.MapClaims)
//
//			// JWT Token 中缺少 sub
//			if claims["sub"] == nil {
//				ctx.Abort(401, "401")
//			}
//
//			// sub 中的值和 userId 不一致，鉴权不通过
//			sub := fmt.Sprintf("%v", claims["sub"])
//			if userId != sub {
//				ctx.Abort(401, "401")
//			}
//		}
//	}
//
//	if err != nil || !token.Valid {
//		ctx.RenderMethodResult("adasd")
//		return
//		//ctx.Abort(401, "401")
//	}
//
//	global.UserId = userId
//}
