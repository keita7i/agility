package spec

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/onsi/ginkgo"
)

var defaultEnvironment = map[string]string{
	"PORT":              "9090",
	"JIRA_API_ENDPOINT": "http://test:9999",
	"JIRA_USER_NAME":    "",
	"JIRA_PASSWORD":     "",
	"REDIS_ADDRS":       "localhost:6379",
	"REDIS_PASSWORD":    "",
	"TEAM_BOARD_IDS":    "test:1000",
}

func StartAgility(environment map[string]string) (*os.Process, error) {
	cmd := exec.Command("go", "run", "github.com/keitam0/agility")
	cmd.Stdout = ginkgo.GinkgoWriter
	cmd.Stderr = ginkgo.GinkgoWriter
	env := map[string]string{}
	for k, v := range defaultEnvironment {
		env[k] = v
	}
	for k, v := range environment {
		env[k] = v
	}
	cmd.Env = append([]string{}, os.Environ()...)
	for k, v := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	for wait := 1; wait <= 32; wait *= 2 {
		af := time.After(time.Duration(wait) * time.Second)
		conn, err := net.DialTimeout("tcp", "localhost:9090", time.Duration(wait)*time.Second)
		if err != nil {
			<-af
			continue // timeout or connection refused
		}
		conn.Close()
		return cmd.Process, nil
	}
	return nil, errors.New("agility failed to start")
}
