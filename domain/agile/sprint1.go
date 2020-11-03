package agile

import "strconv"

type Sprint1 struct {
	name   string
	done   bool
	issues []Issue
}

func NewSprint1(name string) Sprint1 {
	return Sprint1{
		name:   name,
		issues: []Issue{},
		done:   false,
	}
}

func (s Sprint1) Name() string {
	return s.name
}

func (s Sprint1) Done() bool {
	return s.done
}

func (s *Sprint1) SetDone(done bool) {
	s.done = done
}

func (s *Sprint1) AddIssue(i Issue) {
	s.issues = append(s.issues, i)
}

func (s Sprint1) Commitment() int {
	v := 0
	for _, i := range s.issues {
		v += i.Size()
	}
	return v
}

func (s Sprint1) Velocity() int {
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

func (s Sprint1) IsStale() bool {
	return !SprintNameRegex.MatchString(s.Name())
}

func (s Sprint1) SprintNumber() int {
	m := SprintNameRegex.FindStringSubmatch(s.Name())
	if len(m) != 2 {
		return -1
	}
	n, _ := strconv.Atoi(m[1])
	return n
}

func (s Sprint1) Less(other Sprint1) bool {
	return s.SprintNumber() < other.SprintNumber()
}
