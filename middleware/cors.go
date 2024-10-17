package middleware

import (
	"github.com/gin-contrib/cors"
	"time"
)

func NewCors() {
	cors.New(cors.Config{
		AllowOriginFunc:        func(origin string) bool { return true },
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowPrivateNetwork:    true,
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"Origin", "Content-Length", "Content-Type"},
		MaxAge:                 12 * time.Hour,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	})
}
