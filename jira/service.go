package jira

import (
	"regexp"
	"sort"
	"strconv"

	"github.com/keitam913/agility/agile"
)

var SprintNameRegex = regexp.MustCompile(`^[Ss](\d+)$`)

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

	var sps []agile.Sprint
	for _, rs := range rss[from:to] {
		sp, err := s.sprint(rs.Name, rs.State == SprintClosed)
		if err != nil {
			return nil, err
		}
		sps = append(sps, sp)
	}
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
		css := append([]Sprint{}, i.Fields.ClosedSprints...)
		sort.Slice(css, func(i, j int) bool {
			return s.CompareSprint(css[i], css[j]) > 0
		})
		closedSprint := ""
		if len(css) > 0 {
			closedSprint = css[0].Name
		}
		sp.AddIssue(agile.NewIssue(int(i.Fields.Size), i.Fields.Labels, closedSprint))
	}
	return sp, nil
}

func (s *Service) CompareSprint(l, r Sprint) int {
	if !SprintNameRegex.MatchString(l.Name) {
		return -1
	}
	if !SprintNameRegex.MatchString(r.Name) {
		return 1
	}
	ls, err := strconv.Atoi(SprintNameRegex.FindStringSubmatch(l.Name)[1])
	if err != nil {
		panic(err)
	}
	rs, err := strconv.Atoi(SprintNameRegex.FindStringSubmatch(r.Name)[1])
	if err != nil {
		panic(err)
	}
	return ls - rs
}
