package middleware

import (
	"github.com/gin-gonic/gin"
	"hello/service"
)

func LoadMiddlewares(router *gin.Engine) {

	router.Use(injectData) //middleware for inject data

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(accessLog(service.HttpLogger))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(recoveryWithLog(service.ErrorLogger, true))

	service.InitLogger.Info("load all middleware successful")
}
