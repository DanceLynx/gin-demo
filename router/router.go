package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello/controller"
	"hello/middleware"
)

func LoadRoutes(router *gin.Engine) {

	gin.DisableConsoleColor()
	//404错误
	router.NoRoute(controller.NoRoute)
	//405 method not exist
	router.NoMethod(controller.NoMethod)
	//jwt login
	router.GET("/login", controller.Login)

	v1 := router.Group("/v1")
	v1.Use(middleware.AuthRequired)
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.GET("/home", controller.Index)
		v1.GET("/testredis", controller.TestRedis)
		v1.GET("/testdb", controller.TestDB)
		v1.GET("/test", controller.Test)
		v1.GET("/query", controller.TestQuery)
		v1.GET("/bind", controller.TestBind)
		v1.GET("/userinfo", controller.UserInfo)
	}

	fmt.Println("load routes successful.")
}
