package main

import (
	"fmt"
	"github.com/Fordisk123/ginframe/conf"
	"github.com/Fordisk123/ginframe/errors"
	"github.com/Fordisk123/ginframe/frame"
	"github.com/Fordisk123/ginframe/log"
	"github.com/Fordisk123/ginframe/response"
	"github.com/gin-gonic/gin"
)

func main() {

	conf.InitConf("./conf")

	log.NewDefaultLogger("example", "v1.0.0")

	frame.GinServe(Router, true)

}

func Router(r *gin.Engine) {

	r.GET("/ok", func(c *gin.Context) {

		log.WithFields(c, "ok", "ok")

		log.Infof(c, "11111")

		log.WithFields(c, "ok2", "ok2")

		log.Infof(c, "2222")

		response.JsonResponse(c, "ok")
	})
	r.GET("/error400", func(c *gin.Context) {
		response.JsonResponse(c, errors.NewBadRequestError("input error", fmt.Errorf("detail error")))
	})
	r.GET("/error500", func(c *gin.Context) {
		response.JsonResponse(c, errors.NewInternalServerError("internal error", fmt.Errorf("detail error")))
	})
}
