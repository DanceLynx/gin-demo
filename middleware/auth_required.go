package middleware

import (
	"github.com/gin-gonic/gin"
	"hello/config"
	"hello/constant"
	"hello/controller"
	"hello/service"
	"hello/util"
	"strconv"
)

func AuthRequired(ctx *gin.Context) {
	authToken := ctx.Request.Header.Get("Authentication")
	if authToken == "" {
		controller.StatusUnauthorized(ctx)
		return
	}
	uid, err := util.ParseToken(authToken, config.App.JWT_TOKEN)
	if err != nil {
		controller.Error(ctx, constant.USER_JWT_PARSE_FAILD, "非法登录", gin.H{})
		service.Logger.Error(ctx,"登录校验",err)
		return
	}
	//校验redis中是否存在
	val, _ := service.Redis.Exists(ctx, "jwt:user:"+uid).Result()
	if val<=0 {
		controller.Error(ctx, constant.REDIS_KEY_NOT_EXISTS_ERR, "token已过期,请重新登录", gin.H{})
		return
	}

	//刷新token
	Iuid, err := strconv.Atoi(uid)
	if err!=nil {
		panic(err)
	}
	access_token, err := util.CreateToken(uint(Iuid), config.App.JWT_TOKEN)
	if err != nil {
		controller.Error(ctx, constant.REDIS_ERROR,err.Error(), gin.H{})
		return
	}
	ctx.Writer.Header().Set("Authentication", access_token)


	ctx.Set("userId", uid)
	ctx.Set("authToken", authToken)
	ctx.Next()
}
