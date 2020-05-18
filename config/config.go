package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Teams           []string
	JIRAAPIEndpoint string
	JIRAUsername    string
	JIRAPassword    string
	JIRABoardID     string
	RedisAddr       string
	RedisPassword   string
	RedisDB         int
}

func FromEnv() (Config, error) {
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	return Config{
		Teams:           strings.Split(os.Getenv("TEAMS"), ","),
		JIRAAPIEndpoint: os.Getenv("JIRA_API_ENDPOINT"),
		JIRAUsername:    os.Getenv("JIRA_USERNAME"),
		JIRAPassword:    os.Getenv("JIRA_PASSWORD"),
		JIRABoardID:     os.Getenv("JIRA_BOARD_ID"),
		RedisAddr:       os.Getenv("REDIS_ADDR"),
		RedisPassword:   os.Getenv("REDIS_PASSWORD"),
		RedisDB:         redisDB,
	}, nil
}
