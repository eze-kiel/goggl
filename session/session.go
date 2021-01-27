package session

type Session struct {
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Duration  string `json:"duration"`
}

func New() *Session {
	s := Session{}
	return &s
}
