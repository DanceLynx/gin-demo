package middleware

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"hello/service"
	"time"
)

func LoadMiddlewares(router *gin.Engine) {

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(service.HttpLogger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(service.ErrorLogger, true))

	service.InitLogger.Info("middleware", "load all middleware successful")
}
