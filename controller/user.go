package controller

import (
	"github.com/gin-gonic/gin"
	"hello/config"
	"hello/constant"
	"hello/model"
	"hello/service"
	"hello/util"
	"strconv"
)

func Login(ctx *gin.Context) {

	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		Error(ctx, constant.USER_LOGIN_FAILED, "登录失败", gin.H{})
		return
	}
	dbUser, err := service.User.FindByNameAndPass(user.Username, user.Password)
	if err != nil {
		Error(ctx, constant.USER_NOT_EXISTS, "用户不存在", gin.H{})
		return
	}
	access_token, err := util.CreateToken(dbUser.ID, config.App.JWT_TOKEN)
	if err != nil {
		Error(ctx, constant.USER_JWT_ERROR, "登录失败", gin.H{})
		return
	}
	err = service.Redis.HMSet(ctx, "jwt:user:"+strconv.Itoa(int(dbUser.ID)), 
		map[string]interface{}{
			"userId":dbUser.ID,
			"username":dbUser.Username,
		},
	).Err()

	if err != nil {
		Error(ctx, constant.REDIS_ERROR,err.Error(), gin.H{})
		return
	}
	ctx.Writer.Header().Set("Authentication", access_token)
	Success(ctx, "登录成功", gin.H{})

}

func UserInfo(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	Success(ctx, "用户信息如下", gin.H{"userId": userId})
}
