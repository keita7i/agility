package config

import (
	"os"
	"strings"
)

type Config struct {
	Teams           []string
	JIRAAPIEndpoint string
	JIRAUsername    string
	JIRAPassword    string
	JIRABoardID     string
	RedisAddrs      string
	RedisPassword   string
}

func FromEnv() (Config, error) {
	return Config{
		Teams:           strings.Split(os.Getenv("TEAMS"), ","),
		JIRAAPIEndpoint: os.Getenv("JIRA_API_ENDPOINT"),
		JIRAUsername:    os.Getenv("JIRA_USERNAME"),
		JIRAPassword:    os.Getenv("JIRA_PASSWORD"),
		JIRABoardID:     os.Getenv("JIRA_BOARD_ID"),
		RedisAddrs:      os.Getenv("REDIS_ADDRS"),
		RedisPassword:   os.Getenv("REDIS_PASSWORD"),
	}, nil
}
