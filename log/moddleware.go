package log

import (
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		log := WithFields(ctx, "path", c.Request.URL.Path, "method", c.Request.Method)
		logCtx := WithContext(ctx, log)
		c.Request = c.Request.WithContext(logCtx)
		c.Next()
	}
}
