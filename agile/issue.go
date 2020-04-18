package agile

const (
	StatusInProgress = iota + 1
	StatusDone
)

type Issue struct {
	size   int
	labels []string
	status int
}

func NewIssue(size int, labels []string, status int) Issue {
	return Issue{
		size:   size,
		labels: labels,
	}
}

func (i Issue) Size() int {
	return i.size
}

func (i Issue) HasDone() bool {
	return i.status == StatusDone
}

func (i Issue) HasCommittedBy(team string) bool {
	for _, l := range i.labels {
		if l == team {
			return true
		}
	}
	return false
}
