package jira_test

import (
	"sync"
	"testing"

	"github.com/keitam913/agility/config"
	"github.com/keitam913/agility/jira"
)

func BenchmarkGetRawSprints(b *testing.B) {
	s := makeService()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.GetRawSprints()
	}
}

func BenchmarkLastSprints(b *testing.B) {
	s := makeService()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.LastSprints(1)
	}
}

func makeService() *jira.Service {
	conf, err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	return &jira.Service{
		APIEndpoint: conf.JIRAAPIEndpoint,
		Username:    conf.JIRAUsername,
		Password:    conf.JIRAPassword,
		BoardID:     conf.JIRABoardID,
		SprintCache: &sync.Map{},
	}
}
