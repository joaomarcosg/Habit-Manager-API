package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
)

type MockHabitStore struct{}

func (m *MockHabitStore) CreateHabit(
	ctx context.Context,
	name,
	category,
	description string,
	frequency []store.WeekDay,
	startDate,
	targetDate time.Time,
	priority int,
) (uuid.UUID, error) {
	id := uuid.New()
	return id, nil
}

func (m *MockHabitStore) GetHabitById(ctx context.Context, id uuid.UUID) (store.Habit, error) {
	return store.Habit{
		ID:          id,
		Name:        "Work out",
		Category:    "Health",
		Description: "Work out 5 times a week",
		Frequency:   []store.WeekDay{"monday", "tuesday", "wednesday", "thursday", "friday"},
		StartDate:   time.Now(),
		TargetDate:  time.Now().Add(7),
		Priority:    10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockHabitStore) GetHabitByName(ctx context.Context, name string) (store.Habit, error) {
	id := uuid.New()
	return store.Habit{
		ID:          id,
		Name:        name,
		Category:    "Health",
		Description: "Work out 5 times a week",
		Frequency:   []store.WeekDay{"monday", "tuesday", "wednesday", "thursday", "friday"},
		StartDate:   time.Now(),
		TargetDate:  time.Now().Add(7),
		Priority:    10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockHabitStore) UpdateHabit(
	ctx context.Context,
	name,
	category,
	description string,
	frequency []store.WeekDay,
	startDate,
	targetDate time.Time,
	priority int,
) (store.Habit, error) {
	id := uuid.New()
	return store.Habit{
		ID:          id,
		Name:        name,
		Category:    category,
		Description: description,
		Frequency:   frequency,
		StartDate:   startDate,
		TargetDate:  targetDate,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockHabitStore) DeleteHabit(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	return id, nil
}
