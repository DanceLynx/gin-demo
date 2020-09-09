package middleware

import (
	"github.com/gin-gonic/gin"
	"hello/controller"
	"hello/util"
	"hello/config"
	"hello/constant"
)

func AuthRequired(ctx *gin.Context) {
	authToken := ctx.Request.Header.Get("Authentication")
	if authToken == "" {
		controller.StatusUnauthorized(ctx)
		ctx.Abort()
	}
	uid,err := util.ParseToken(authToken,config.App.JWT_TOKEN)
	if err!=nil {
		controller.Error(ctx,constant.USER_VERIFY_FAILD,"认证失败",gin.H{})
	}
	ctx.Set("userId",uid)
	ctx.Set("authToken",authToken)
	ctx.Next()
}