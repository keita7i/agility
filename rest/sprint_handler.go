package rest

import (
	"net/http"

	"github.com/keitam913/agility/application"

	"github.com/gin-gonic/gin"
)

type SprintHandler struct {
	ApplicationService *application.Service
	Teams              []string
}

func (sh *SprintHandler) GET(ctx *gin.Context) {
	sps, err := sh.ApplicationService.LastSprints(2)
	if err != nil {
		panic(err)
	}
	srs := make([]SprintResponse, 0)
	for i, sp := range sps {
		sr := SprintResponse{}
		sr.Sprint = sp.Sprint()
		sr.Teams = make(map[string]TeamMetrics, 0)
		sr.Teams["All"] = TeamMetrics{
			Commitment: sp.AllCommitment(),
			Done:       sp.AllDone(),
			Velocity:   sp.AllVelocity(sps[i+1:]),
		}
		for _, t := range sh.Teams {
			sr.Teams[t] = TeamMetrics{
				Commitment: sp.Commitment(t),
				Done:       sp.Done(t),
				Velocity:   sp.Velocity(t, sps[i+1:]),
			}
		}
		srs = append(srs, sr)
	}
	ctx.JSON(http.StatusOK, srs)
}
