package redis

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/go-redis/redis/v7"
	"github.com/keitam0/agility/infra/jira"
)

const (
	SprintKey  = "sprints"
	IssueKey   = "issues/%s"
	SprintBKey = "sprints/%s"
	IssueBKey  = "issues/%s/%s"
)

var (
	SprintCacheExpiration = 1 * time.Minute
)

type CachedJIRAClient struct {
	JIRAClient  jira.Client
	RedisClient Client
}

func (c *CachedJIRAClient) Sprints(boardID string) ([]jira.Sprint, error) {
	var sprints []jira.Sprint
	b, err := c.RedisClient.Get(fmt.Sprintf(SprintKey, boardID)).Bytes()
	if err == redis.Nil {
		ss, err := c.JIRAClient.Sprints(boardID)
		if err != nil {
			return nil, err
		}
		b, err := json.Marshal(ss)
		if err != nil {
			return nil, err
		}
		if err := c.RedisClient.Set(fmt.Sprintf(SprintKey, boardID), b, SprintCacheExpiration).Err(); err != nil {
			return nil, err
		}
		sprints = ss
	} else if err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(b, &sprints); err != nil {
			return nil, err
		}
	}
	return sprints, nil
}

func (c *CachedJIRAClient) Issues(boardID string, sprint string, sprintDone bool) ([]jira.Issue, error) {
	var issues []jira.Issue
	b, err := c.RedisClient.Get(fmt.Sprintf(IssueBKey, boardID, sprint)).Bytes()
	if err == redis.Nil {
		is, err := c.JIRAClient.Issues(boardID, sprint, sprintDone)
		if err != nil {
			return nil, err
		}
		b, err := json.Marshal(is)
		if err != nil {
			return nil, err
		}
		if sprintDone {
			if err := c.RedisClient.Set(fmt.Sprintf(IssueBKey, boardID, sprint), b, 0).Err(); err != nil {
				return nil, err
			}
		}
		issues = is
	} else if err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(b, &issues); err != nil {
			return nil, err
		}
	}
	return issues, nil
}
