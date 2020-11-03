package jira

import (
	"context"
	"fmt"
	"sort"
	"strconv"

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
		return s.CompareSprint(rss[i], rss[j]) > 0
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
			sp, err := s.sprint(s.TeamBoardIDs[team], rst.Name, rst.State == SprintClosed)
			if err != nil {
				return err
			}
			if !sp.IsStale() {
				spCh <- sp
			}
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

func (s *Service) sprint(boardID, sprint string, done bool) (agile.Sprint, error) {
	sp := agile.NewSprint(sprint)
	sp.SetDone(done)
	is, err := s.Client.Issues(boardID, sprint, done)
	if err != nil {
		return agile.Sprint{}, err
	}
	for _, i := range is {
		var assignedSprints []Sprint
		if i.Fields.Sprint.Name != "" {
			assignedSprints = append(assignedSprints, i.Fields.Sprint)
		}
		assignedSprints = append(assignedSprints, i.Fields.ClosedSprints...)
		closedSprint := ""
		if isDoneStatus(i.Fields.Status.Name) && len(assignedSprints) > 0 {
			latestSprint := assignedSprints[0]
			for j := 1; j < len(i.Fields.ClosedSprints); j++ {
				if s.CompareSprint(latestSprint, assignedSprints[j]) < 0 {
					latestSprint = assignedSprints[j]
				}
			}
			closedSprint = latestSprint.Name
		}
		sp.AddIssue(agile.NewIssue(int(i.Fields.Size), i.Fields.Labels, closedSprint))
	}
	return sp, nil
}

func (s *Service) CompareSprint(l, r Sprint) int {
	if !agile.SprintNameRegex.MatchString(l.Name) {
		return -1
	}
	if !agile.SprintNameRegex.MatchString(r.Name) {
		return 1
	}
	ls, err := strconv.Atoi(agile.SprintNameRegex.FindStringSubmatch(l.Name)[1])
	if err != nil {
		panic(err)
	}
	rs, err := strconv.Atoi(agile.SprintNameRegex.FindStringSubmatch(r.Name)[1])
	if err != nil {
		panic(err)
	}
	return ls - rs
}

func isDoneStatus(statusName string) bool {
	switch statusName {
	case "完了", "クローズ", "解決済み":
		return true
	default:
		return false
	}
}
