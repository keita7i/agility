package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/keitam0/agility/application"
)

type BoardHandler struct {
	ApplicationService *application.Service
}

func (bh *BoardHandler) GET(ctx *gin.Context) {

}
