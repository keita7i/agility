package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keitam913/agility/config"
	"github.com/keitam913/agility/rest"
)

type DI struct {
}

func (di DI) Router() *gin.Engine {
	return rest.NewRouter()
}

func (DI) Config() config.Config {
	c, err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	return c
}
