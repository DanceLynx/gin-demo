package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"strings"
)

func injectData(ctx *gin.Context) {

	traceId := ksuid.New().String()

	ctx.Set("traceId", strings.ToLower(traceId))
	ctx.Next()
}
