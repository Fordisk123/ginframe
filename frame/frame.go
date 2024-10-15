package frame

import (
	"github.com/Fordisk123/ginframe/log"
	gsession "github.com/Fordisk123/ginframe/pkg/session"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/ory/viper"
	"github.com/penglongli/gin-metrics/ginmetrics"
)
import "github.com/gin-gonic/gin"

type Route func(router *gin.Engine)

func GinServe(r Route, block bool, cMiddleware ...gin.HandlerFunc) {

	router := gin.New()

	m := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 0.5, 1, 3, 5, 10})

	router.Use(log.GetGinLogger(), gin.Recovery(), log.LoggerMiddleware(), sessions.Sessions("mmpsession", gsession.SessionStore))
	m.Use(router)
	if cMiddleware != nil && len(cMiddleware) > 0 {
		router.Use(cMiddleware...)
	}

	gin.SetMode(gin.DebugMode)

	//http端口默认8080
	httpPort := viper.GetString("http.port")
	if httpPort == "" {
		httpPort = "8080"
	}

	addr := "0.0.0.0:" + httpPort

	r(router)

	pprof.Register(router)

	if block {
		if err := router.Run(addr); err != nil {
			panic("启动gin失败:" + err.Error())
		}
	} else {
		go func() {
			if err := router.Run(addr); err != nil {
				panic("启动gin失败:" + err.Error())
			}
		}()
	}
}
