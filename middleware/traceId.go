package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"time"
)

func injectTraceId(ctx *gin.Context) {
	timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)

	requestId := ksuid.New().String()
	context.WithValue(timeoutContext, "traceId", requestId)
	ctx.Set("traceId", requestId)
	ctx.Next()
}
