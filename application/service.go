package application

import (
	"errors"
	"strings"
	"sync"

	"github.com/keitam0/agility/agile"
)

const MAX_SPRINTS = 8

type Service struct {
	JIRAService JIRAService
	Teams       []string
}

func (s Service) LastSprints(max int) ([]agile.Sprint, error) {
	return s.JIRAService.LastSprints(max)
}

func (s Service) BoardOfTeam(team string) (agile.Board, error) {
	return s.JIRAService.BoardOfTeam(team, MAX_SPRINTS)
}

func (s Service) AllBoards() ([]agile.Board, error) {
	errCh := make(chan error, len(s.Teams))
	defer close(errCh)
	sm := &sync.Map{}
	wg := &sync.WaitGroup{}
	for _, t := range s.Teams {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			b, err := s.BoardOfTeam(t)
			if err != nil {
				errCh <- err
				return
			}
			sm.Store(t, b)
		}(t)
	}
	wg.Wait()
	if len(errCh) > 0 {
		var errs []string
		for err := range errCh {
			errs = append(errs, err.Error())
		}
		return nil, errors.New(strings.Join(errs, ","))
	}
	var bs []agile.Board
	for _, t := range s.Teams {
		b, ok := sm.Load(t)
		if !ok {
			panic("must not happen")
		}
		bs = append(bs, b.(agile.Board))
	}
	return bs, nil
}
