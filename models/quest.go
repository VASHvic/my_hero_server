package models

// Quest represents a quest that a user can complete
type Quest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Stat        string `json:"stat"`
}
