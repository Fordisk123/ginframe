package jwt

import (
	"fmt"
	"net/http"
	"strings"
)

const JwtHttpRequestHeaderKey = "Authorization"

func ValidHttpRequestWithJwt(c *http.Request, jwt Jwt) (map[string]interface{}, error) {
	token := c.Header.Get(JwtHttpRequestHeaderKey)
	if token == "" {
		return nil, fmt.Errorf("missing token")
	}
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimLeft(token, "Bearer")
	}
	token = strings.TrimSpace(token)
	return jwt.Parse(token)
}

func SetToken(resp *http.Response, jwt Jwt, payLoad map[string]interface{}) error {
	tokenStr, err := jwt.Generate(payLoad)
	if err != nil {
		return err
	}
	resp.Header.Set(JwtHttpRequestHeaderKey, tokenStr)
	return nil
}
