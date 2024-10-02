package pkg

import (
	"net/http"
	"strings"
)

func GetHttpRequestClientIp(r http.Request) string {
	header := r.Header.Get("X-Forwarded-For")
	if header != "" {
		s := strings.Split(header, ",")
		return s[0]
	}
	ip := r.Header.Get("X-Real-Ip")
	if ip != "" {
		return ip
	}
	return r.RemoteAddr
}
