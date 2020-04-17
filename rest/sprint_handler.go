package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SprintHandler struct {
}

func (sh *SprintHandler) GET(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []SprintResponse{
		{
			Sprint: "s12",
			Teams: map[string]TeamMetrics{
				"All": {
					Commitment: 10,
					Done: 10,
					Velocity: 10,
				},
				"SRE 0": {
					Commitment: 10,
					Done: 10,
					Velocity: 10,
				},
				"SRE 1+2": {
					Commitment: 10,
					Done: 10,
					Velocity: 10,
				},
			},
		},
	})
}
