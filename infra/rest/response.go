package rest

type SprintResponse struct {
	Sprint string                 `json:"sprint"`
	Teams  map[string]TeamMetrics `json:"teams"`
}

type TeamMetrics struct {
	Commitment int `json:"commitment"`
	Done       int `json:"done"`
	Velocity   int `json:"velocity"`
}

type BoardResponse struct {
	Team                      string   `json:"team"`
	Sprints                   []Sprint `json:"sprints"`
	AverageOfLatestVelocities int      `json:"averageOfLatestVelocities"`
}

type Sprint struct {
	Name       string `json:"name"`
	Commitment int    `json:"commitment"`
	Velocity   int    `json:"velocity"`
}
