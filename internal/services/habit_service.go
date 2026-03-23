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
