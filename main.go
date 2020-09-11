package main

import (
	"github.com/gin-gonic/gin"
	"hello/config"
	"hello/middleware"
	"hello/router"
	"hello/service"
)

func init() {
	config.LoadConfig() //加载配置
}

func main() {
	r := gin.New()

	service.InitLog()      //配置日志
	service.ConnectRedis() //连接redis
	service.ConnectDB()    //连接数据库
	service.AutoMigrate()  //自动迁移

	middleware.LoadMiddlewares(r) //加载中间件

	router.LoadRoutes(r)   //加载路由
	r.Run(config.App.Port) // listen and serve on 0.0.0.0:8080

	resourceRelease()
}

func resourceRelease() {
	go func() {
		service.DisconnectDB()
	}()
}
