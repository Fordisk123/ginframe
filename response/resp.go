package response

import (
	"github.com/Fordisk123/ginframe/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	RtnCode string
	Data    interface{}
}

func JsonResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, response{
		RtnCode: "000000",
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, err errors.RequestError) {
	c.AbortWithStatusJSON(http.StatusOK, err)
}
