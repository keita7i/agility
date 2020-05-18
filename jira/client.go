package jira

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	APIEndpoint string
	Username    string
	Password    string
	BoardID     string
}

func (c *Client) Sprints() ([]Sprint, error) {
	var ss []Sprint
	maxResults := 0
	isLast := false
	for startAt := 0; !isLast; startAt += maxResults {
		u, err := url.Parse(fmt.Sprintf("%s/agile/1.0/board/%s/sprint", c.APIEndpoint, c.BoardID))
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
		req.SetBasicAuth(c.Username, c.Password)
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
		ss = append(ss, sr.Values...)

		maxResults = int(sr.MaxResults)
		isLast = sr.IsLast
	}
	return ss, nil
}

func (c *Client) Issues(sprint string) ([]Issue, error) {
	var issues []Issue
	total := math.MaxInt32
	maxResults := 0
	for startAt := 0; startAt < total; startAt += maxResults {
		u, err := url.Parse(fmt.Sprintf("%s/agile/1.0/board/%s/issue", c.APIEndpoint, c.BoardID))
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
		req.SetBasicAuth(c.Username, c.Password)

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
