package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keitam0/agility/application"
)

type BoardHandler struct {
	ApplicationService *application.Service
}

func (bh *BoardHandler) GET(ctx *gin.Context) {
	b, err := bh.ApplicationService.BoardOfTeam(ctx.Param("team"))
	if err != nil {
		panic(err)
	}
	ss := make([]Sprint, 0)
	for _, s := range b.Sprints() {
		ss = append(ss, Sprint{
			Name:       s.Name(),
			Commitment: s.Commitment(),
			Velocity:   s.Velocity(),
		})
	}
	ctx.JSON(http.StatusOK, BoardResponse{
		Team:                      b.Team(),
		Sprints:                   ss,
		AverageOfLatestVelocities: b.AverageOfVelocityOfLastThreeSprints(),
	})
}
