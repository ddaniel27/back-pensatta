package domain

type Exercise struct {
	ID    uint64  `json:"id"`
	Score float64 `json:"score"`
	Time  uint64  `json:"time"`
}
