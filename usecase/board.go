package usecase

import (
	"errors"
	"strings"
	"sync"

	"github.com/keitam0/agility/domain/agile"
)

const MAX_SPRINTS = 8

type Board struct {
	JIRAService JIRAService
	Teams       []string
}

func (b Board) LastSprints(max int) ([]agile.Sprint, error) {
	return b.JIRAService.LastSprints(max)
}

func (b Board) BoardOfTeam(team string) (agile.Board, error) {
	return b.JIRAService.BoardOfTeam(team, MAX_SPRINTS)
}

func (b Board) AllBoards() ([]agile.Board, error) {
	errCh := make(chan error, len(b.Teams))
	defer close(errCh)
	sm := &sync.Map{}
	wg := &sync.WaitGroup{}
	for _, t := range b.Teams {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			b, err := b.BoardOfTeam(t)
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
	for _, t := range b.Teams {
		b, ok := sm.Load(t)
		if !ok {
			panic("must not happen")
		}
		bs = append(bs, b.(agile.Board))
	}
	return bs, nil
}
