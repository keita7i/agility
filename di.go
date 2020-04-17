package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keitam913/agility/config"
	"github.com/keitam913/agility/rest"
)

type DI struct {
}

func (di DI) Router() *gin.Engine {
	return rest.NewRouter(di.SprintHandler())
}

func (di DI) SprintHandler() *rest.SprintHandler {
	return &rest.SprintHandler{}
}

func (DI) Config() config.Config {
	c, err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	return c
}
