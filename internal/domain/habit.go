package domain

import (
	"time"

	"github.com/google/uuid"
)

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
	ID          uuid.UUID
	Name        string
	Category    string
	Description string
	Frequency   []WeekDay
	StartDate   time.Time
	TargetDate  time.Time
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
