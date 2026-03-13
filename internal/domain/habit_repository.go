package domain

import (
	"context"

	"github.com/google/uuid"
)

type HabitRepository interface {
	CreateHabit(ctx context.Context, habit Habit) (uuid.UUID, error)
	GetHabitByName(ctx context.Context, name string) (Habit, error)
	UpdateHabit(ctx context.Context, habit Habit) (Habit, error)
	DeleteHabit(ctx context.Context, name string) error
}
