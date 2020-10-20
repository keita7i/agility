package agile

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
	from := len(b.sprints) - 3
	if from < 0 {
		from = 0
	}
	total := 0
	for i := from; i < len(b.sprints); i++ {
		total += b.sprints[i].Velocity()
	}
	return total / (len(b.sprints) - from)
}
