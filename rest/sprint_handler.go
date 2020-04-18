package rest

import (
	"ghe.corp.yahoo.co.jp/pivotal-cf/agility2/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SprintHandler struct {
	ApplicationService application.Service
}

func (sh *SprintHandler) GET(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []SprintResponse{
		{
			Sprint: 12,
			Teams: map[string]TeamMetrics{
				"All": {
					Commitment: 10,
					Done: 10,
					Velocity: 10,
				},
				"SRE0": {
					Commitment: 10,
					Done: 10,
					Velocity: 10,
				},
				"SRE1+2": {
					Commitment: 10,
					Done: 10,
					Velocity: 10,
				},
			},
		},
	})
}
