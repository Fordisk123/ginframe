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
	err := session.Save()
	if err != nil {
		response.ErrorResponse(c, errors.NewInternalServerError("", err))
		return
	}
	err = captcha.WriteImage(c.Writer, id, captcha.StdWidth, captcha.StdHeight)
	if err != nil {
		response.ErrorResponse(c, errors.NewBadRequestError("", err))
		return
	}
	c.Status(http.StatusOK)
	return
}

var Verify = func(c *gin.Context, id string) error {
	session := sessions.Default(c)
	if session == nil {
		return errors.NewBadRequestError("", fmt.Errorf("session is not valid"))
	}
	gc := session.Get("captcha")
	if gc == nil {
		return errors.NewBadRequestError("", fmt.Errorf("captcha is not valid"))
	}
	if !captcha.VerifyString(gc.(string), id) {
		return errors.NewBadRequestError("", fmt.Errorf("captcha is not valid"))
	}
	return nil
}

var Clean = func(c *gin.Context) error {
	session := sessions.Default(c)
	session.Delete("captcha")
	session.Save()
	return nil
}
