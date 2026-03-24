package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

type MockCategoryRepo struct {
	CreateCategoryFn    func(ctx context.Context, category domain.Category) (uuid.UUID, error)
	GetCategoryByNameFn func(ctx context.Context, name string) (domain.Category, error)
	IncrementEntriesFn  func(ctx context.Context, name string) error
	DeleteCategoryFn    func(ctx context.Context, name string) error
}

func (m *MockCategoryRepo) CreateCategory(ctx context.Context, category domain.Category) (uuid.UUID, error) {
	return m.CreateCategoryFn(ctx, category)
}

func (m *MockCategoryRepo) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {
	return m.GetCategoryByNameFn(ctx, name)
}

func (m *MockCategoryRepo) IncrementEntries(ctx context.Context, name string) error {
	return m.IncrementEntriesFn(ctx, name)
}

func (m *MockCategoryRepo) DeleteCategory(ctx context.Context, name string) error {
	return m.DeleteCategoryFn(ctx, name)
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

func (m *MockHabitRepository) CreateHabitWithCategoryUpdate(
	ctx context.Context,
	habit domain.Habit,
) (uuid.UUID, error) {
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

func TestCreateHabit_Success(t *testing.T) {
	expectedID := uuid.New()
	called := false

	mockHabitRepo := &MockHabitRepository{
		CreateHabitFn: func(ctx context.Context, habit domain.Habit) (uuid.UUID, error) {
			called = true
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

	if !called {
		t.Fatalf("expected CreateHabitWithCategoryUpdate to be called")
	}
}

func TestCreateHabit_CategoryNotFound(t *testing.T) {

	mockHabitRepo := &MockHabitRepository{
		CreateHabitFn: func(ctx context.Context, habit domain.Habit) (uuid.UUID, error) {
			return uuid.New(), domain.ErrCategoryNotFound
		},
	}

	mockCategoryRepo := &MockCategoryRepo{
		GetCategoryByNameFn: func(ctx context.Context, name string) (domain.Category, error) {
			return domain.Category{}, domain.ErrCategoryNotFound
		},
	}

	service := NewHabitService(mockHabitRepo, mockCategoryRepo)

	newHabit := domain.Habit{
		Name:     "Work out",
		Category: "Health",
	}

	id, err := service.CreateHabit(context.Background(), newHabit)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, domain.ErrCategoryNotFound) {
		t.Fatalf("expected ErrCategoryNotFound, got %v", err)
	}

	if id != uuid.Nil {
		t.Fatalf("expected empty uuid, got %v", id)
	}

}

func TestGetHabitByName_Success(t *testing.T) {
	expectedHabit := domain.Habit{
		ID:          uuid.New(),
		Name:        "Work out",
		Category:    "Health",
		Description: "Work out five days a week",
		Frequency:   []domain.WeekDay{"monday", "tuesday", "wednesday", "thursday", "friday"},
		StartDate:   time.Now(),
		TargetDate:  time.Now().Add(7),
		Priority:    10,
	}

	mockHabitRepo := &MockHabitRepository{
		GetHabitByNameFn: func(ctx context.Context, name string) (domain.Habit, error) {
			return domain.Habit{
				ID:          expectedHabit.ID,
				Name:        expectedHabit.Name,
				Category:    expectedHabit.Category,
				Description: expectedHabit.Description,
				Frequency:   expectedHabit.Frequency,
				StartDate:   expectedHabit.StartDate,
				TargetDate:  expectedHabit.TargetDate,
				Priority:    expectedHabit.Priority,
			}, nil
		},
	}

	mockCategoryRepo := &MockCategoryRepo{}

	service := NewHabitService(mockHabitRepo, mockCategoryRepo)

	habit, err := service.GetHabitByName(context.Background(), "Work out")

	if err != nil {
		t.Fatalf("unexpected erro %v", err)
	}

	if habit != expectedHabit {
		t.Fatalf("expected %v, got %v", expectedHabit, habit)
	}
}
