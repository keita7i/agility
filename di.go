package main

import (
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/keitam0/agility/infra/config"
	"github.com/keitam0/agility/infra/jira"
	agredis "github.com/keitam0/agility/infra/redis"
	"github.com/keitam0/agility/infra/rest"
	"github.com/keitam0/agility/usecase"
)

type DI struct {
}

func (di DI) Router() *gin.Engine {
	return rest.NewRouter(di.BoardHandler())
}

func (di DI) BoardHandler() *rest.BoardHandler {
	return &rest.BoardHandler{
		Board: di.BoardUsecase(),
	}
}

func (di DI) BoardUsecase() *usecase.Board {
	ids := di.Config().TeamBoardIDs0
	var teams []string
	for _, tbi := range ids {
		teams = append(teams, tbi.Team)
	}
	return &usecase.Board{
		JIRAService: di.JIRAService(),
		Teams:       teams,
	}
}

func (di DI) JIRAService() usecase.JIRAService {
	return &jira.Service{
		Client:       di.JIRAClient(),
		TeamBoardIDs: di.Config().TeamBoardIDs,
	}
}

func (di DI) JIRAClient() jira.Client {
	conf := di.Config()
	return &agredis.CachedJIRAClient{
		JIRAClient:  jira.NewClient(conf.JIRAAPIEndpoint, conf.JIRAUsername, conf.JIRAPassword),
		RedisClient: di.RedisClient(),
	}
}

func (di DI) RedisClient() agredis.Client {
	conf := di.Config()
	addrs := strings.Split(conf.RedisAddrs, ",")
	if len(addrs) > 1 {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        addrs,
			Password:     conf.RedisPassword,
			MaxRedirects: 1,
			MaxRetries:   0,
			DialTimeout:  1000 * time.Millisecond,
			ReadTimeout:  1000 * time.Millisecond,
			PoolSize:     16 * runtime.NumCPU(),
			MinIdleConns: 0,
			PoolTimeout:  100 * time.Millisecond,
		})
	}
	return redis.NewClient(&redis.Options{
		Addr:     addrs[0],
		Password: conf.RedisPassword,
	})
}

func (DI) Config() config.Config {
	c, err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	return c
}
