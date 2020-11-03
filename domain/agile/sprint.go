package agile

import (
	"regexp"
	"strconv"
)

var SprintNameRegex = regexp.MustCompile(`^[Pp](\d+).*$`)

type Sprint struct {
	name   string
	done   bool
	issues []Issue
}

func NewSprint(name string) Sprint {
	return Sprint{
		name:   name,
		issues: []Issue{},
		done:   false,
	}
}

func (s Sprint) Name() string {
	return s.name
}

func (s Sprint) Done() bool {
	return s.done
}

func (s *Sprint) SetDone(done bool) {
	s.done = done
}

func (s *Sprint) AddIssue(i Issue) {
	s.issues = append(s.issues, i)
}

func (s Sprint) Commitment() int {
	v := 0
	for _, i := range s.issues {
		v += i.Size()
	}
	return v
}

func (s Sprint) Velocity() int {
	if !s.done {
		return -1
	}
	v := 0
	for _, i := range s.issues {
		if i.DoneSprint() == s.Name() {
			v += i.Size()
		}
	}
	return v
}

func (s Sprint) IsStale() bool {
	return !SprintNameRegex.MatchString(s.Name())
}

func (s Sprint) SprintNumber() int {
	m := SprintNameRegex.FindStringSubmatch(s.Name())
	if len(m) != 2 {
		return -1
	}
	n, _ := strconv.Atoi(m[1])
	return n
}

func (s Sprint) Less(other Sprint) bool {
	return s.SprintNumber() < other.SprintNumber()
}
