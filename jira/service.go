package jira

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/keitam913/agility/agile"
)

var SprintNameRegex = regexp.MustCompile(`^[Ss](\d+)$`)

type Service struct {
	APIEndpoint string
	Username    string
	Password    string
	BoardID     string
	SprintCache *sync.Map
}

func (s *Service) LastSprints(max int) ([]agile.Sprint, error) {
	rss, err := s.GetRawSprints()
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
		sp, err := s.Sprint(rs.Name, rs.State == SprintClosed)
		if err != nil {
			return nil, err
		}
		sps = append(sps, sp)
	}
	return sps, nil
}

func (s *Service) GetRawSprints() ([]Sprint, error) {
	var ss []Sprint
	maxResults := 0
	isLast := false
	for startAt := 0; !isLast; startAt += maxResults {
		u, err := url.Parse(fmt.Sprintf("%s/agile/1.0/board/%s/sprint", s.APIEndpoint, s.BoardID))
		if err != nil {
			return nil, err
		}
		q := u.Query()
		q.Add("startAt", strconv.Itoa(startAt))
		u.RawQuery = q.Encode()
		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			return nil, err
		}
		req.SetBasicAuth(s.Username, s.Password)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			b := &strings.Builder{}
			fmt.Fprintf(b, "status = %s; body = ", res.Status)
			io.Copy(b, res.Body)
			return nil, errors.New(b.String())
		}

		var sr SprintResponse
		if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
			return nil, err
		}
		for _, sv := range sr.Values {
			if SprintNameRegex.MatchString(sv.Name) {
				ss = append(ss, sv)
			}
		}

		maxResults = int(sr.MaxResults)
		isLast = sr.IsLast
	}
	return ss, nil
}

func (s *Service) Sprint(sprint string, done bool) (agile.Sprint, error) {
	si, ok := s.SprintCache.Load(sprint)
	if ok {
		return si.(agile.Sprint), nil
	}

	sp := agile.NewSprint(sprint)
	sp.SetDone(done)

	is, err := s.GetIssues(sprint)
	if err != nil {
		return agile.Sprint{}, err
	}
	for _, i := range is {
		st := agile.StatusInProgress
		if i.Fields.Status.Name == "完了" {
			st = agile.StatusDone
		}
		sp.AddIssue(agile.NewIssue(int(i.Fields.Size), i.Fields.Labels, st))
	}
	if sp.HasClosed() {
		s.SprintCache.Store(sprint, sp)
	}

	return sp, nil
}

func (s *Service) GetIssues(sprint string) ([]Issue, error) {
	var issues []Issue
	total := math.MaxInt32
	maxResults := 0
	for startAt := 0; startAt < total; startAt += maxResults {
		u, err := url.Parse(fmt.Sprintf("%s/agile/1.0/board/%s/issue", s.APIEndpoint, s.BoardID))
		if err != nil {
			return nil, err
		}
		q := u.Query()
		q.Add("jql", fmt.Sprintf("スプリント = \"%s\" AND type IN (\"task\", \"ストーリー\")", sprint))
		q.Add("startAt", strconv.Itoa(startAt))
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			return nil, err
		}
		req.SetBasicAuth(s.Username, s.Password)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			sb := &strings.Builder{}
			io.Copy(sb, res.Body)
			return nil, fmt.Errorf("failed to search issues: Status: %s; Body: %s", res.Status, sb.String())
		}

		var ires IssueResponse
		if err := json.NewDecoder(res.Body).Decode(&ires); err != nil {
			return nil, err
		}

		issues = append(issues, ires.Issues...)

		total = int(ires.Total)
		maxResults = int(ires.MaxResults)
	}
	return issues, nil
}

func (s *Service) CompareSprint(l, r Sprint) int {
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
