package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(sh *SprintHandler, th *TeamHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/v1/sprints", sh.GET)
	r.GET("/v1/teams", th.GET)

	r.StaticFile("/main.js", "/usr/share/agility/assets/main.js")
	r.StaticFile("/main.css", "/usr/share/agility/assets/main.css")

	r.NoRoute(func(ctx *gin.Context) {
		ctx.File("/usr/share/agility/assets/index.html")
	})

	return r
}
