package controller

import (
	"github.com/gin-gonic/gin"
	"hello/model"
	"hello/service"
	"hello/constant"
	"hello/util"
	"hello/config"
	"fmt"
)


func Login(ctx *gin.Context) {

	var user model.User
	if err:=ctx.ShouldBind(&user);err!=nil {
		Error(ctx,constant.USER_LOGIN_FAILED,"登录校验失败",gin.H{})
		return
	}
	dbUser,err := service.User.FindByNameAndPass(user.Username,user.Password)
	if err != nil {
		Error(ctx,constant.USER_NOT_EXISTS,"用户不存在",gin.H{})
		return
	}
	fmt.Println(dbUser)
	fmt.Println(config.App.JWT_TOKEN)
	access_token,err := util.CreateToken(dbUser.ID,config.App.JWT_TOKEN)
	if err!=nil {
		Error(ctx,constant.USER_JWT_ERROR,"登录失败",gin.H{})
		return
	}
	Success(ctx,"登录成功",gin.H{"access_token":access_token})

}

func UserInfo(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	Success(ctx,"用户信息如下",gin.H{"userId":userId})
}