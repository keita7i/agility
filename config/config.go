package config

import (
	"errors"
	"os"
	"strings"
)

type Config struct {
	Teams           []string
	JIRAAPIEndpoint string
	JIRAUsername    string
	JIRAPassword    string
	JIRABoardID     string
	TeamBoardIDs    map[string]string
	RedisAddrs      string
	RedisPassword   string
}

func FromEnv() (Config, error) {
	bIDs := map[string]string{}
	for _, teamToBID := range strings.Split(os.Getenv("TEAM_BOARD_IDS"), ",") {
		p := strings.Split(teamToBID, ":")
		if len(p) != 2 {
			return Config{}, errors.New("TEAM_BOARD_IDS must be specified as a list of pairs")
		}
		bIDs[p[0]] = p[1]
	}
	return Config{
		Teams:           strings.Split(os.Getenv("TEAMS"), ","),
		JIRAAPIEndpoint: os.Getenv("JIRA_API_ENDPOINT"),
		JIRAUsername:    os.Getenv("JIRA_USERNAME"),
		JIRAPassword:    os.Getenv("JIRA_PASSWORD"),
		JIRABoardID:     os.Getenv("JIRA_BOARD_ID"),
		TeamBoardIDs:    bIDs,
		RedisAddrs:      os.Getenv("REDIS_ADDRS"),
		RedisPassword:   os.Getenv("REDIS_PASSWORD"),
	}, nil
}
