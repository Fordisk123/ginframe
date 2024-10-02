package middleware

import (
	"github.com/Fordisk123/ginframe/pkg"
	"github.com/gin-gonic/gin"
)

const ClientIpKey = "mmpClinetIp"

func ClientIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		//st := time.Now()
		defer func() {
			//ft := time.Since(t)
		}()
		c.Set(ClientIpKey, pkg.GetHttpRequestClientIp(*c.Request))
	}
}
