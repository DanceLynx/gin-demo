package middleware

import (
	"hello/core/log"

	"github.com/gin-gonic/gin"
)

func flushLog(ctx *gin.Context) {
	log.LogSync()
}
