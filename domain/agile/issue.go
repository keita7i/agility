package agile

type Issue struct {
	size    int
	labels  []string
	status  string
	sprints []string
}

func NewIssue(size int, labels []string) Issue {
	return Issue{
		size:   size,
		labels: labels,
	}
}

func (i Issue) Size() int {
	return i.size
}

func (i Issue) HasDone() bool {
	switch i.status {
	case "完了", "クローズ", "解決済み":
		return true
	default:
		return false
	}
}

func (i Issue) DoneSprint() string {
	if !i.HasDone() {
		return ""
	}
	if len(i.sprints) <= 0 {
		return ""
	}
	lastSprint := NewSprint(i.sprints[0])
	for j := 1; j < len(i.sprints); j++ {
		sp := NewSprint(i.sprints[j])
		if lastSprint.Less(sp) {
			lastSprint = sp
		}
	}
	return lastSprint.Name()
}

func (i Issue) HasCommittedBy(team string) bool {
	for _, l := range i.labels {
		if l == team {
			return true
		}
	}
	return false
}

func (i *Issue) SetStatus(status string) {
	i.status = status
}

func (i *Issue) AddSprint(sprint string) {
	i.sprints = append(i.sprints, sprint)
}
