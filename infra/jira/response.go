package jira

const (
	SprintClosed = "closed"
)

type IssueResponse struct {
	StartAt    float32 `json:"startAt"`
	MaxResults float32 `json:"maxResults"`
	Total      float32 `json:"total"`
	Issues     []Issue `json:"issues"`
}

type Issue struct {
	Fields Fields `json:"fields"`
}

type Fields struct {
	Summary       string   `json:"summary"`
	Status        Status   `json:"status"`
	Labels        []string `json:"labels"`
	Size          float32  `json:"customfield_10002"`
	Sprint        Sprint   `json:"sprint"`
	ClosedSprints []Sprint `json:"closedSprints"`
}

type Status struct {
	Name string `json:"name"`
}

type SprintResponse struct {
	StartAt    float32  `json:"startAt"`
	MaxResults float32  `json:"maxResults"`
	IsLast     bool     `json:"isLast"`
	Values     []Sprint `json:"values"`
}

type Sprint struct {
	Name      string `json:"name"`
	State     string `json:"state"`
	StartDate string `json:"startDate"`
}
