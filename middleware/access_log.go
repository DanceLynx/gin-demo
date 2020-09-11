package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func accessLog(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Request.ParseForm()
		traceId, _ := c.Get("traceId")
		logger.Info("request",
			zap.String("traceId", traceId.(string)),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),
			zap.String("path", path),
			zap.String("query", query),
			zap.Any("post", c.Request.PostForm),

			//zap.String("user-agent", c.Request.UserAgent()),
		)

		c.Next()
		end := time.Now()
		cost := end.Sub(start)

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			response, _ := c.Get("response")
			logger.Info("response",
				zap.String("traceId", traceId.(string)),
				zap.Int("status", c.Writer.Status()),
				zap.Any("data", response),
				zap.Duration("cost", cost),
			)
		}
	}
}
