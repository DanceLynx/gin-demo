package middleware

import (
	"github.com/gin-gonic/gin"
	"hello/config"
	"hello/constant"
	"hello/controller"
	"hello/util"
)

func AuthRequired(ctx *gin.Context) {
	authToken := ctx.Request.Header.Get("Authentication")
	if authToken == "" {
		controller.StatusUnauthorized(ctx)
		//阻止下面的执行需要用abort & return
		ctx.Abort()
		return
	}
	uid, err := util.ParseToken(authToken, config.App.JWT_TOKEN)
	if err != nil {
		controller.Error(ctx, constant.USER_JWT_PARSE_FAILD, "认证失败", gin.H{})
		ctx.Abort()
		return
	}

	ctx.Set("userId", uid)
	ctx.Set("authToken", authToken)
	ctx.Next()
}
