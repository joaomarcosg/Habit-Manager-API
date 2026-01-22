package store

import "time"

type WeekDay string

const (
	Monday    WeekDay = "monday"
	Tuesday   WeekDay = "tuesday"
	Wednesday WeekDay = "wednesday"
	Thursday  WeekDay = "thursday"
	Friday    WeekDay = "friday"
	Saturday  WeekDay = "saturday"
	Sunday    WeekDay = "sunday"
)

type Habit struct {
	ID          int       `json:"id"`
	Name        string    `json:"habit_name"`
	Category    Category  `json:"habit_category"`
	Description string    `json:"description"`
	Frequency   []WeekDay `json:"frequency"`
	StartDate   time.Time `json:"start_date"`
	TargetDate  time.Time `json:"target_date"`
	Priority    int       `json:"priority"`
}
