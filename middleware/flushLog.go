package middleware

import (
	"fmt"
	"hello/core/log"

	"github.com/gin-gonic/gin"
)

func flushLog(ctx *gin.Context) {
	ctx.Next()
	fmt.Print("flush log")
	log.LogSync()
}
