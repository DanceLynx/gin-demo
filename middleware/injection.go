package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"time"
)

func injectData(ctx *gin.Context) {
	timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)

	requestId := ksuid.New().String()
	context.WithValue(timeoutContext, "requestId", requestId)
	ctx.Set("requestId", requestId)
	ctx.Next()
}
