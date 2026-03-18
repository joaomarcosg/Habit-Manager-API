package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

type MockCategoryRepo struct {
	GetCategoryByNameFn func(ctx context.Context, name string) (domain.Category, error)
}

func (m *MockCategoryRepo) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {
	return m.GetCategoryByNameFn(ctx, name)
}

type MockHabitRepository struct {
	CreateHabitFn    func(ctx context.Context, habit domain.Habit) (uuid.UUID, error)
	GetHabitByNameFn func(ctx context.Context, name string) (domain.Habit, error)
	UpdateHabitFn    func(ctx context.Context, habit domain.Habit) (domain.Habit, error)
	DeleteHabitFn    func(ctx context.Context, name string) error
}

func (m *MockHabitRepository) CreateHabit(ctx context.Context, habit domain.Habit) (uuid.UUID, error) {
	return m.CreateHabitFn(ctx, habit)
}

func (m *MockHabitRepository) GetHabitByName(ctx context.Context, name string) (domain.Habit, error) {
	return m.GetHabitByNameFn(ctx, name)
}

func (m *MockHabitRepository) UpdateHabit(ctx context.Context, habit domain.Habit) (domain.Habit, error) {
	return m.UpdateHabitFn(ctx, habit)
}

func (m *MockHabitRepository) DeleteHabit(ctx context.Context, name string) error {
	return m.DeleteHabitFn(ctx, name)
}

func TestSuccessCreateHabit(t *testing.T) {
	expectedID := uuid.New()

	mockHabitRepo := &MockHabitRepository{
		CreateHabitFn: func(ctx context.Context, habit domain.Habit) (uuid.UUID, error) {
			return expectedID, nil
		},
	}

	mockCategoryRepo := &MockCategoryRepo{
		GetCategoryByNameFn: func(ctx context.Context, name string) (domain.Category, error) {
			return domain.Category{Name: "Health"}, nil
		},
	}

	service := NewHabitService(mockHabitRepo, mockCategoryRepo)

	newHabit := domain.Habit{
		Name:     "Work out",
		Category: "Health",
	}

	id, err := service.CreateHabit(context.Background(), newHabit)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if id != expectedID {
		t.Fatalf("expected %v, got %v", expectedID, id)
	}
}
