package jira

import (
	"context"
	"fmt"
	"sort"

	"github.com/keitam0/agility/domain/agile"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	Client       Client
	TeamBoardIDs map[string]string
}

func (s *Service) BoardOfTeam(team string, maxSprints int) (agile.Board, error) {
	bID, ok := s.TeamBoardIDs[team]
	if !ok {
		return agile.Board{}, fmt.Errorf("invalid team: %s", team)
	}
	rss, err := s.Client.Sprints(bID)
	if err != nil {
		return agile.Board{}, err
	}
	sort.Slice(rss, func(i, j int) bool {
		return agile.NewSprint(rss[j].Name).Less(agile.NewSprint(rss[i].Name))
	})
	from := len(rss)
	for i, rs := range rss {
		if rs.State == SprintClosed {
			from = i - 1
			break
		}
	}
	if from < 0 {
		from = 0
	}
	to := from + maxSprints
	if to > len(rss) {
		to = len(rss)
	}
	spCh := make(chan agile.Sprint, maxSprints)
	eg, _ := errgroup.WithContext(context.TODO())
	for _, rs := range rss[from:to] {
		rst := rs
		eg.Go(func() error {
			sp := agile.NewSprint(rst.Name)
			if sp.IsStale() {
				return nil
			}
			iss, err := s.issues(s.TeamBoardIDs[team], rst.Name)
			if err != nil {
				return err
			}
			for _, is := range iss {
				sp.AddIssue(is)
			}
			sp.SetDone(rst.State == SprintClosed)
			spCh <- sp
			return nil
		})
	}
	err = eg.Wait()
	close(spCh)
	if err != nil {
		return agile.Board{}, err
	}
	var sps []agile.Sprint
	for sp := range spCh {
		sps = append(sps, sp)
	}
	sort.Slice(sps, func(i, j int) bool {
		return sps[j].Less(sps[i])
	})

	b := agile.NewBoard(team)
	for _, sp := range sps {
		b.AddSprint(sp)
	}

	return b, nil
}

func (s *Service) issues(boardID, sprint string) ([]agile.Issue, error) {
	irs, err := s.Client.Issues(boardID, sprint, true)
	if err != nil {
		return nil, err
	}
	iss := []agile.Issue{}
	for _, ir := range irs {
		is := agile.NewIssue(int(ir.Fields.Size), ir.Fields.Labels)
		is.SetStatus(ir.Fields.Status.Name)
		if ir.Fields.Sprint.Name != "" {
			is.AddSprint(ir.Fields.Sprint.Name)
		}
		for _, sp := range ir.Fields.ClosedSprints {
			is.AddSprint(sp.Name)
		}
		iss = append(iss, is)
	}
	return iss, nil
}
