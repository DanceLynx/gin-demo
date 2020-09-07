package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello/model"
	"hello/service"
	"time"
)

func Index(ctx *gin.Context) {

	Success(ctx, "获取成功", gin.H{"hello": "test"})
}

func TestRedis(ctx *gin.Context) {
	err := service.Redis.Set(ctx, "name", "Hello JSON", 5*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	val, err := service.Redis.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	Success(ctx, "获取成功", gin.H{"test": val})
}

func TestDB(ctx *gin.Context) {
	user := model.User{
		Username: "范兄弟",
		Password: "3333",
	}
	service.User.AddUser(&user)
	fmt.Println(user)

	user1, err1 := service.User.GetUserById(1)
	if err1 != nil {
		service.Logger.Error("user1 error", err1)
	}
	service.Logger.Info("user1", user1)
	Success(ctx, "成功", gin.H{"data": user1})
}

func Test(ctx *gin.Context) {
	panic("this is me")
}
