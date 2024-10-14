package captcha

import (
	"fmt"
	"github.com/Fordisk123/ginframe/errors"
	"github.com/Fordisk123/ginframe/response"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Captcha = func(c *gin.Context, length int) {
	id := captcha.NewLen(length)
	session := sessions.Default(c)
	session.Set("captcha", id)
	err := captcha.WriteImage(c.Writer, id, captcha.StdWidth, captcha.StdHeight)
	if err != nil {
		response.ErrorResponse(c, errors.NewBadRequestError("", err))
		return
	}
	c.Status(http.StatusOK)
	return
}

var Verify = func(c *gin.Context, id string) error {
	gc := sessions.Default(c).Get("captcha")
	if gc != id {
		return errors.NewBadRequestError("", fmt.Errorf("captcha is not valid"))
	}
	return nil
}
