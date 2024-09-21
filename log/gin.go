package log

import "github.com/gin-gonic/gin"

func GetGinLogger() gin.HandlerFunc {
	return gin.LoggerWithWriter(DefaultLogger.GetWriter())
}
