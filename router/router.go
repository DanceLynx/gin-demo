package router

import (
	"github.com/gin-gonic/gin"
	"hello/controller"
	"fmt"
)

func LoadRoutes(router *gin.Engine) {

	//404错误
	router.NoRoute(controller.NoRoute)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/home", controller.Index)
	router.GET("/testredis", controller.TestRedis)
	router.GET("/testdb", controller.TestDB)
	router.GET("/test", controller.Test)

	fmt.Println("load routes successful.")
}
