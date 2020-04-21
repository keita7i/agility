package agile

import (
	"math"
)

type Sprint struct {
	sprint int
	issues []Issue
	done   bool
}

func NewSprint(sprint int) Sprint {
	return Sprint{
		sprint: sprint,
		issues: []Issue{},
		done:   false,
	}
}

func (s Sprint) Sprint() int {
	return s.sprint
}

func (s *Sprint) AddIssue(issue Issue) {
	s.issues = append(s.issues, issue)
}

func (s Sprint) Issues() []Issue {
	return s.issues
}

func (s *Sprint) SetDone(done bool) {
	s.done = done
}

func (s Sprint) AllCommitment() int {
	c := 0
	for _, i := range s.issues {
		c += i.Size()
	}
	return c
}

func (s Sprint) AllDone() int {
	c := 0
	for _, i := range s.issues {
		if i.HasDone() {
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
	c := 0
	for _, i := range s.issues {
		if i.HasCommittedBy(team) && i.HasDone() {
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
