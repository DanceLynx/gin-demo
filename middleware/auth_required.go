package middleware

import (
	"hello/config"
	"hello/constant"
	"hello/controller"
	"hello/core/log"
	"hello/core/redis"
	"hello/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthRequired(ctx *gin.Context) {
	authToken := ctx.Request.Header.Get("Authentication")
	if authToken == "" {
		controller.ErrorWithMessage(ctx, constant.USER_JWT_PARSE_FAILD, "登录失败")
		ctx.Abort()
		return
	}
	uid, err := util.ParseToken(authToken, config.App.JWT_TOKEN)
	if err != nil {
		controller.Error(ctx, constant.USER_JWT_PARSE_FAILD)
		log.Error(ctx, "auth parse jwt", err)
		ctx.Abort()
		return
	}
	//校验redis中是否存在
	val, err := redis.Client.Exists(ctx, "jwt:user:"+uid).Result()
	if err != nil {
		log.Error(ctx, "auth redis", err)
		controller.Error(ctx, constant.REDIS_KEY_NOT_EXISTS_ERR)
		ctx.Abort()
		return
	}
	if val <= 0 {
		controller.Error(ctx, constant.REDIS_KEY_NOT_EXISTS_ERR)
		ctx.Abort()
		return
	}
	//刷新token
	Iuid, err := strconv.Atoi(uid)
	if err != nil {
		panic(err)
	}
	accessToken, err := util.CreateToken(uint(Iuid), config.App.JWT_TOKEN)
	if err != nil {
		controller.ErrorWithMessage(ctx, constant.REDIS_ERROR, err.Error())
		ctx.Abort()
		return
	}
	ctx.Writer.Header().Set("Authentication", accessToken)

	ctx.Set("userId", uid)
	ctx.Set("authToken", authToken)
}
