package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(sh *SprintHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/v1/sprints", sh.GET)

	r.StaticFile("/main.js", "/usr/share/agility/assets/main.js")
	r.StaticFile("/main.css", "/usr/share/agility/assets/main.css")

	r.NoRoute(func(ctx *gin.Context) {
		ctx.File("/usr/share/agility/assets/index.html")
	})

	return r
}
