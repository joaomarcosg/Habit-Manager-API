package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

type HabitService struct {
	repo         domain.HabitRepository
	categoryRepo domain.CategoryRepository
}

func NewHabitService(
	repo domain.HabitRepository,
	categoryRepo domain.CategoryRepository,
) *HabitService {
	return &HabitService{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (hs *HabitService) CreateHabit(ctx context.Context, habit domain.Habit) (uuid.UUID, error) {

	id, err := hs.repo.CreateHabit(ctx, habit)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil

}

func (hs *HabitService) GetHabitByName(ctx context.Context, name string) (domain.Habit, error) {

	habit, err := hs.repo.GetHabitByName(ctx, name)

	if err != nil {
		return domain.Habit{}, err
	}

	return domain.Habit{
		ID:          habit.ID,
		Name:        habit.Name,
		Category:    habit.Category,
		Description: habit.Description,
		Frequency:   habit.Frequency,
		StartDate:   habit.StartDate,
		TargetDate:  habit.TargetDate,
		Priority:    habit.Priority,
	}, nil

}

func (hs *HabitService) UpdateHabit(ctx context.Context, habit domain.Habit) (domain.Habit, error) {

	habit, err := hs.repo.UpdateHabit(ctx, habit)

	if err != nil {
		return domain.Habit{}, err
	}

	return habit, nil

}
