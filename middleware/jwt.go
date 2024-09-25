package middleware

import (
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
