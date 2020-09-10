package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"time"
)

func injectData(ctx *gin.Context) {
	timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)

	traceId := ksuid.New().String()
	context.WithValue(timeoutContext, "traceId", traceId)
	ctx.Set("traceId", traceId)
	ctx.Next()
}