package jira

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/keitam0/agility/agile"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	Client       Client
	TeamBoardIDs map[string]string
}

func (s *Service) LastSprints(max int) ([]agile.Sprint, error) {
	rss, err := s.Client.Sprints()
	if err != nil {
		return nil, err
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
	to := from + max
	if to > len(rss) {
		to = len(rss)
	}
	spCh := make(chan agile.Sprint, max)
	eg, _ := errgroup.WithContext(context.TODO())
	for _, rs := range rss[from:to] {
		rst := rs
		eg.Go(func() error {
			sp, err := s.sprint(rst.Name, rst.State == SprintClosed)
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
		return nil, err
	}
	var sps []agile.Sprint
	for sp := range spCh {
		sps = append(sps, sp)
	}
	sort.Slice(sps, func(i, j int) bool {
		return sps[j].Less(sps[i])
	})

	return sps, nil
}

func (s *Service) BoardOfTeam(team string, maxSprints int) (agile.Board, error) {
	bID, ok := s.TeamBoardIDs[team]
	if !ok {
		return agile.Board{}, fmt.Errorf("invalid team: %s", team)
	}
	rss, err := s.Client.SprintsB(bID)
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
	spCh := make(chan agile.Sprint1, maxSprints)
	eg, _ := errgroup.WithContext(context.TODO())
	for _, rs := range rss[from:to] {
		rst := rs
		eg.Go(func() error {
			sp, err := s.sprintB(s.TeamBoardIDs[team], rst.Name, rst.State == SprintClosed)
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
	var sps []agile.Sprint1
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

func (s *Service) sprint(sprint string, done bool) (agile.Sprint, error) {
	sp := agile.NewSprint(sprint)
	sp.SetDone(done)
	is, err := s.Client.Issues(sprint, done)
	if err != nil {
		return agile.Sprint{}, err
	}
	for _, i := range is {
		closedSprint := ""
		if isDoneStatus(i.Fields.Status.Name) && len(i.Fields.ClosedSprints) > 0 {
			lastClosedSprint := i.Fields.ClosedSprints[0]
			for j := 1; j < len(i.Fields.ClosedSprints); j++ {
				if s.CompareSprint(lastClosedSprint, i.Fields.ClosedSprints[j]) < 0 {
					lastClosedSprint = i.Fields.ClosedSprints[j]
				}
			}
			closedSprint = lastClosedSprint.Name
		}
		sp.AddIssue(agile.NewIssue(int(i.Fields.Size), i.Fields.Labels, closedSprint))
	}
	return sp, nil
}

func (s *Service) sprintB(boardID, sprint string, done bool) (agile.Sprint1, error) {
	sp := agile.NewSprint1(sprint)
	sp.SetDone(done)
	is, err := s.Client.IssuesB(boardID, sprint, done)
	if err != nil {
		return agile.Sprint1{}, err
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
