package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
	"runtime/debug"
	"net"
	"strings"
	"time"
	"fmt"
	"hello/controller"
)

func LoadMiddlewares(router *gin.Engine) {

	router.Use(CaptureRequestAndResponse)

	router.Use(Recovery(true))
	fmt.Println("load all middleware successful.")
}

func CaptureRequestAndResponse(ctx *gin.Context) {
	//
	ctx.Next()
	value,exists := ctx.Get("response")
	if exists {
		ctx.JSON(ctx.Writer.Status(),value)
	}
}

func  Recovery(stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				fmt.Println("error",err)
				if brokenPipe {
					
					// If the connection is dead, we can't write a status to it.
					ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				}

				fmt.Println("time",time.Now())
				
				if stack {
					fmt.Println("stack",string(debug.Stack()))
				}

				controller.StatusInternalServerError(ctx)
			}
		}()
		ctx.Next()
	}
}