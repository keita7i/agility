package redis

import (
	"github.com/keitam913/agility/jira"
)

type CachedJIRAClient struct {
	JIRAClient jira.Client
}

func (c *CachedJIRAClient) Sprints() ([]jira.Sprint, error) {
	return c.JIRAClient.Sprints()
}

func (c *CachedJIRAClient) Issues(sprint string) ([]jira.Issue, error) {
	return c.Issues(sprint)
}
