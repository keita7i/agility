package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keitam0/agility/usecase"
)

type BoardHandler struct {
	ApplicationService *usecase.Service
}

func (bh *BoardHandler) GET(ctx *gin.Context) {
	bs, err := bh.ApplicationService.AllBoards()
	if err != nil {
		panic(err)
	}
	brs := make([]BoardResponse, 0)
	for _, b := range bs {
		ss := make([]Sprint, 0)
		for _, s := range b.Sprints() {
			ss = append(ss, Sprint{
				Name:       s.Name(),
				Commitment: s.Commitment(),
				Velocity:   s.Velocity(),
			})
		}
		brs = append(brs, BoardResponse{
			Team:                      b.Team(),
			Sprints:                   ss,
			AverageOfLatestVelocities: b.AverageOfVelocityOfLastThreeSprints(),
		})
	}
	ctx.JSON(http.StatusOK, brs)
}
