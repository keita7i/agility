package rest

import (
	"net/http"

	"github.com/keitam913/agility/application"

	"github.com/gin-gonic/gin"
)

const SHOW_SPRINTS = 8

type SprintHandler struct {
	ApplicationService *application.Service
	Teams              []string
}

func (sh *SprintHandler) GET(ctx *gin.Context) {
	sps, err := sh.ApplicationService.LastSprints(SHOW_SPRINTS + 2) // Get more sprints to calculate velocities
	if err != nil {
		panic(err)
	}
	srs := make([]SprintResponse, 0)
	for i := 0; i < len(sps) && i < SHOW_SPRINTS; i++ {
		sr := SprintResponse{}
		sr.Sprint = sps[i].Sprint()
		sr.Teams = make(map[string]TeamMetrics, 0)
		sr.Teams["All"] = TeamMetrics{
			Commitment: sps[i].AllCommitment(),
			Done:       sps[i].AllDone(),
			Velocity:   sps[i].AllVelocity(sps[i+1:]),
		}
		for _, t := range sh.Teams {
			sr.Teams[t] = TeamMetrics{
				Commitment: sps[i].Commitment(t),
				Done:       sps[i].Done(t),
				Velocity:   sps[i].Velocity(t, sps[i+1:]),
			}
		}
		srs = append(srs, sr)
	}
	ctx.JSON(http.StatusOK, srs)
}
