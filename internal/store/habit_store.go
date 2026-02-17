package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

type HabitStore interface {
	CreateHabit(
		ctx context.Context,
		name,
		category,
		description string,
		frequency []domain.WeekDay,
		startDate,
		targetDate time.Time,
		priority int,
	) (uuid.UUID, error)
	GetHabitById(ctx context.Context, id uuid.UUID) (domain.Habit, error)
	GetHabitByName(ctx context.Context, name string) (domain.Habit, error)
	UpdateHabit(
		ctx context.Context,
		name,
		category,
		description string,
		frequency []domain.WeekDay,
		startDate,
		targetDate time.Time,
	) (domain.Habit, error)
	DeleteHabit(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}
