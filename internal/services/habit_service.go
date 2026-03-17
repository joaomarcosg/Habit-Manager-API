package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

type HabitService struct {
	repo domain.HabitRepository
}

func NewHabitService(repo domain.HabitRepository) *HabitService {
	return &HabitService{
		repo: repo,
	}
}

func (hs *HabitService) CreateHabit(ctx context.Context, habit domain.Habit) (uuid.UUID, error) {

	newHabit := domain.Habit{
		Name:        habit.Name,
		Category:    habit.Category,
		Description: habit.Description,
		Frequency:   habit.Frequency,
		StartDate:   habit.StartDate,
		TargetDate:  habit.TargetDate,
		Priority:    habit.Priority,
	}

	id, err := hs.repo.CreateHabit(ctx, newHabit)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil

}
