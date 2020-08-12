package jira

import (
	"context"
	"sort"
	"strconv"

	"github.com/keita7i/agility/agile"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	Client Client
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

func (s *Service) sprint(sprint string, done bool) (agile.Sprint, error) {
	sp := agile.NewSprint(sprint)
	sp.SetDone(done)
	is, err := s.Client.Issues(sprint, done)
	if err != nil {
		return agile.Sprint{}, err
	}
	for _, i := range is {
		closedSprint := ""
		if i.Fields.Status.Name == "完了" && len(i.Fields.ClosedSprints) > 0 {
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
