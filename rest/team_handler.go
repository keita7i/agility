package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	Teams []string
}

func (th *TeamHandler) GET(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, th.Teams)
}
