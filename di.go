package main

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/keitam913/agility/application"
	"github.com/keitam913/agility/config"
	"github.com/keitam913/agility/jira"
	"github.com/keitam913/agility/rest"
)

type DI struct {
}

func (di DI) Router() *gin.Engine {
	return rest.NewRouter(di.SprintHandler(), di.TeamHandler())
}

func (di DI) SprintHandler() *rest.SprintHandler {
	return &rest.SprintHandler{
		ApplicationService: di.ApplicationService(),
		Teams:              di.Config().Teams,
	}
}

func (di DI) TeamHandler() *rest.TeamHandler {
	return &rest.TeamHandler{
		Teams: di.Config().Teams,
	}
}

func (di DI) ApplicationService() *application.Service {
	return &application.Service{
		JIRAService: di.JIRAService(),
	}
}

func (di DI) JIRAService() application.JIRAService {
	conf := di.Config()
	return &jira.Service{
		APIEndpoint: conf.JIRAAPIEndpoint,
		Username:    conf.JIRAUsername,
		Password:    conf.JIRAPassword,
		BoardID:     conf.JIRABoardID,
		SprintCache: &sync.Map{},
	}
}

func (DI) Config() config.Config {
	c, err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	return c
}
