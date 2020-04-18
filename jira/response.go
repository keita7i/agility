package jira

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
	Status Status   `json:"status"`
	Labels []string `json:"labels"`
	Size   float32  `json:"customfield_10002"`
}

type Status struct {
	Name string `json:"name"`
}
