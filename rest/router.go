package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(sh *SprintHandler, th *TeamHandler, bh *BoardHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/v1/sprints", sh.GET)
	r.GET("/v1/teams", th.GET)
	r.GET("/v1/boards", bh.GET)

	ph := promhttp.Handler()
	r.GET("/metrics", func(ctx *gin.Context) {
		ph.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.StaticFile("/main.js", "/usr/share/agility/assets/main.js")
	r.StaticFile("/main.css", "/usr/share/agility/assets/main.css")

	r.NoRoute(func(ctx *gin.Context) {
		ctx.File("/usr/share/agility/assets/index.html")
	})

	return r
}
