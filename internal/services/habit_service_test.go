package services

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

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

	mockRepo := &MockHabitRepository{
		CreateHabitFn: func(ctx context.Context, habit domain.Habit) (uuid.UUID, error) {
			return expectedID, nil
		},
	}

	service := NewHabitService(mockRepo)

	id, err := service.CreateHabit(
		context.Background(),
		"Work out",
		"Health",
		"Work out 5 times a week",
		[]domain.WeekDay{"monday", "tuesday", "wednesday", "thursday", "friday"},
		time.Now(),
		time.Now().Add(7),
		10,
	)

	if err != nil {
		t.Fatal("unexpected error %v", err)
	}

	if id != expectedID {
		t.Fatalf("expected %v, got %v", expectedID, id)
	}
}
