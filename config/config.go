package config

import (
	"os"
	"strings"
)

type Config struct {
	Teams []string
}

func FromEnv() (Config, error) {
	return Config{
		Teams: strings.Split(os.Getenv("TEAMS"), ","),
	}, nil
}
