package middleware

import (
	"fmt"
	"github.com/Fordisk123/ginframe/errors"
	"github.com/Fordisk123/ginframe/pkg/jwt"
	"github.com/Fordisk123/ginframe/response"
	"github.com/gin-gonic/gin"
)

func JwtMiddleWare(jwter jwt.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		//st := time.Now()
		defer func() {
			//ft := time.Since(t)
		}()
		tokenPayLoad, err := jwt.ValidHttpRequestWithJwt(c.Request, jwter)
		if err != nil {
			response.ErrorResponse(c, errors.NewBadRequestError("", err))
			return
		}
		c.Set("tokenData", tokenPayLoad)
	}
}

func GetJwtPayLoad(c *gin.Context) (map[string]interface{}, error) {
	value, exists := c.Get("tokenData")
	if !exists {
		return nil, fmt.Errorf("can't find token")
	}
	ms, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("%+V token is invalid", value)
	}
	return ms, nil
}
