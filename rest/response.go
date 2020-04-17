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
