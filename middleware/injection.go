package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func injectData(ctx *gin.Context) {

	traceId := ksuid.New().String()

	ctx.Set("traceId", strings.ToLower(traceId))
	fmt.Println("inject data")
}
