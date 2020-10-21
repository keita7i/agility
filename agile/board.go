package agile

import (
	"math"
)

type Board struct {
	team    string
	sprints []Sprint1
}

func NewBoard(team string) Board {
	return Board{
		team:    team,
		sprints: []Sprint1{},
	}
}

func (b Board) Team() string {
	return b.team
}

func (b Board) Sprints() []Sprint1 {
	return b.sprints
}

func (b *Board) AddSprint(sprint Sprint1) {
	b.sprints = append(b.sprints, sprint)
}

func (b *Board) AverageOfVelocityOfLastThreeSprints() int {
	sum := 0
	cnt := 0
	for i := 0; i < len(b.Sprints()) && cnt < 3; i++ {
		if !b.Sprints()[i].Done() {
			continue
		}
		sum += b.Sprints()[i].Velocity()
		cnt++
	}
	return int(math.Round(float64(sum) / float64(cnt)))
}
