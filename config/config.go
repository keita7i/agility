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
	JIRAProjectKey  string
	JIRABoardID     string
}

func FromEnv() (Config, error) {
	return Config{
		Teams:           strings.Split(os.Getenv("TEAMS"), ","),
		JIRAAPIEndpoint: os.Getenv("JIRA_API_ENDPOINT"),
		JIRAUsername:    os.Getenv("JIRA_USERNAME"),
		JIRAPassword:    os.Getenv("JIRA_PASSWORD"),
		JIRAProjectKey:  os.Getenv("JIRA_PROJECT_KEY"),
		JIRABoardID:     os.Getenv("JIRA_BOARD_ID"),
	}, nil
}
