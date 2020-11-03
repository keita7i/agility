package config

import (
	"errors"
	"os"
	"strings"
)

type Config struct {
	JIRAAPIEndpoint string
	JIRAUsername    string
	JIRAPassword    string
	TeamBoardIDs    map[string]string
	TeamBoardIDs0   []TeamBoardID
	RedisAddrs      string
	RedisPassword   string
}

type TeamBoardID struct {
	Team    string
	BoardID string
}

func FromEnv() (Config, error) {
	var bIDs0 []TeamBoardID
	bIDs := map[string]string{}
	for _, teamToBID := range strings.Split(os.Getenv("TEAM_BOARD_IDS"), ",") {
		p := strings.Split(teamToBID, ":")
		if len(p) != 2 {
			return Config{}, errors.New("TEAM_BOARD_IDS must be specified as a list of pairs")
		}
		bIDs[p[0]] = p[1]
		bIDs0 = append(bIDs0, TeamBoardID{Team: p[0], BoardID: p[1]})
	}
	return Config{
		JIRAAPIEndpoint: os.Getenv("JIRA_API_ENDPOINT"),
		JIRAUsername:    os.Getenv("JIRA_USERNAME"),
		JIRAPassword:    os.Getenv("JIRA_PASSWORD"),
		TeamBoardIDs:    bIDs,
		TeamBoardIDs0:   bIDs0,
		RedisAddrs:      os.Getenv("REDIS_ADDRS"),
		RedisPassword:   os.Getenv("REDIS_PASSWORD"),
	}, nil
}
