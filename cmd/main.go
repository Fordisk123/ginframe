package main

import (
	"github.com/Fordisk123/ginframe/conf"
	"github.com/Fordisk123/ginframe/frame"
	"github.com/Fordisk123/ginframe/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	conf.InitConf("./conf")

	log.NewDefaultLogger("example", "v1.0.0")

	frame.GinServe(Router, true)

}

func Router(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}
