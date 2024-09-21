package log

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		WithFields(c, "requestId", uuid.New().String(), "path", c.Request.URL.Path, "method", c.Request.Method)
		c.Next()
	}
}
