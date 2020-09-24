package controller

import (
	"hello/config"
	"hello/constant"
	"hello/core/log"
	"hello/core/redis"
	"hello/model"
	"hello/service"
	"hello/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		ErrorWithMessage(ctx, constant.USER_LOGIN_FAILED, "登录失败")
		return
	}
	err := service.UserService.Find(ctx, &user)
	log.Info(ctx, "user", user)
	if err != nil {
		ErrorWithMessage(ctx, constant.USER_NOT_EXISTS, "用户不存在")
		return
	}
	access_token, err := util.CreateToken(user.ID, config.App.JWT_TOKEN)
	if err != nil {
		Error(ctx, constant.USER_JWT_ERROR)
		return
	}
	err = redis.Client.HMSet(ctx, "jwt:user:"+strconv.Itoa(int(user.ID)),
		map[string]interface{}{
			"userId":   user.ID,
			"username": user.Username,
		},
	).Err()

	if err != nil {
		log.Info(ctx, "login jwt", err)
		ErrorWithMessage(ctx, constant.REDIS_ERROR, err.Error())
		return
	}
	ctx.Writer.Header().Set("Authentication", access_token)
	Success(ctx, "登录成功", gin.H{"info": user})

}

func UserInfo(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	Success(ctx, "用户信息如下", gin.H{"userId": userId})
}
