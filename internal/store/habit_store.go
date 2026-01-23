package store

import (
	"context"
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
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"habit_name"`
	Category    Category  `json:"habit_category"`
	Description string    `json:"description"`
	Frequency   []WeekDay `json:"frequency"`
	StartDate   time.Time `json:"start_date"`
	TargetDate  time.Time `json:"target_date"`
	Priority    int       `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type HabitStore interface {
	CreateHabit(
		ctx context.Context,
		name string,
		category Category,
		description string,
		frequency []WeekDay,
		startDate,
		targetDate time.Time,
		priority int,
	) (uuid.UUID, error)
	GetHabitById(ctx context.Context, id uuid.UUID) (Habit, error)
	GetHabitByName(ctx context.Context, name string) (Habit, error)
	UpdateHabit(
		ctx context.Context,
		name string,
		category Category,
		description string,
		frequency []WeekDay,
		startDate,
		targetDate time.Time,
	) (Habit, error)
	DeleteHabit(ctx context.Context, name string) (bool, error)
}
