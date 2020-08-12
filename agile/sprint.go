package agile

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

var SprintNameRegex = regexp.MustCompile(`^[Pp](\d+).*$`)

type Sprint struct {
	sprint    string
	issues    []Issue
	startedAt time.Time
	done      bool
}

func NewSprint(sprint string) Sprint {
	return Sprint{
		sprint: sprint,
		issues: []Issue{},
		done:   false,
	}
}

func (s Sprint) Sprint() string {
	if s.IsStale() {
		return ""
	}
	n := SprintNameRegex.FindStringSubmatch(s.sprint)[1]
	return fmt.Sprintf("p%s", n)
}

func (s Sprint) IsStale() bool {
	return !SprintNameRegex.MatchString(s.sprint)
}

func (s *Sprint) Start(at time.Time) {
	s.startedAt = at
}

func (s *Sprint) AddIssue(issue Issue) {
	s.issues = append(s.issues, issue)
}

func (s Sprint) StartedAt() time.Time {
	return s.startedAt
}

func (s Sprint) Issues() []Issue {
	return s.issues
}

func (s *Sprint) SetDone(done bool) {
	s.done = done
}

func (s Sprint) HasClosed() bool {
	return s.done
}

func (s Sprint) AllCommitment() int {
	c := 0
	for _, i := range s.issues {
		c += i.Size()
	}
	return c
}

func (s Sprint) AllDone() int {
	if !s.done {
		return -1
	}
	c := 0
	for _, i := range s.issues {
		if i.DoneSprint() == s.sprint {
			c += i.Size()
		}
	}
	return c
}

func (s Sprint) AllVelocity(lastSprints []Sprint) int {
	if !s.done {
		return -1
	}
	sum := s.AllDone()
	for i := 0; i < 2 && i < len(lastSprints); i++ {
		sum += lastSprints[i].AllDone()
	}
	return int(math.Round(float64(sum) / 3))
}

func (s Sprint) Commitment(team string) int {
	c := 0
	for _, i := range s.issues {
		if i.HasCommittedBy(team) {
			c += i.Size()
		}
	}
	return c
}

func (s Sprint) Done(team string) int {
	if !s.done {
		return -1
	}
	c := 0
	for _, i := range s.issues {
		if i.HasCommittedBy(team) && i.DoneSprint() == s.sprint {
			c += i.Size()
		}
	}
	return c
}

func (s Sprint) Velocity(team string, lastSprints []Sprint) int {
	if !s.done {
		return -1
	}
	sum := s.Done(team)
	for i := 0; i < 2 && i < len(lastSprints); i++ {
		sum += lastSprints[i].Done(team)
	}
	return int(math.Round(float64(sum) / 3))
}

func (s Sprint) SprintNumber() int {
	m := SprintNameRegex.FindStringSubmatch(s.sprint)
	if len(m) != 2 {
		return -1
	}
	n, _ := strconv.Atoi(m[1])
	return n
}

func (s Sprint) Less(other Sprint) bool {
	return s.SprintNumber() < other.SprintNumber()
}
