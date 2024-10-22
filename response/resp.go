package response

import (
	"github.com/Fordisk123/ginframe/errors"
	"github.com/Fordisk123/ginframe/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	RtnCode int         `json:"code"`
	Data    interface{} `json:"data"`
}

func JsonResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, response{
		RtnCode: 200,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, err errors.RequestError) {
	c.AbortWithStatusJSON(http.StatusOK, err)
}

func Response(c *gin.Context, contentType string, data []byte) {
	c.Writer.Header().Set("Content-Type", contentType)
	_, err := c.Writer.Write(data)
	if err != nil {
		log.GetLogger(c.Request.Context()).Errorf("write response error: %s", err.Error())
	}
	c.Status(http.StatusOK)
}
