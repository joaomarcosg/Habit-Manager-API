package api

import (
	"context"
	"encoding/gob"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/services"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
)

type MockHabitStore struct{}

// DeleteHabit implements store.HabitStore.
func (m *MockHabitStore) DeleteHabit(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	panic("unimplemented")
}

// GetHabitById implements store.HabitStore.
func (m *MockHabitStore) GetHabitById(ctx context.Context, id uuid.UUID) (store.Habit, error) {
	panic("unimplemented")
}

// GetHabitByName implements store.HabitStore.
func (m *MockHabitStore) GetHabitByName(ctx context.Context, name string) (store.Habit, error) {
	panic("unimplemented")
}

// UpdateHabit implements store.HabitStore.
func (m *MockHabitStore) UpdateHabit(ctx context.Context, name string, category string, description string, frequency []store.WeekDay, startDate time.Time, targetDate time.Time) (store.Habit, error) {
	panic("unimplemented")
}

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
	habit := store.Habit{
		ID:          uuid.New(),
		Name:        name,
		Category:    category,
		Description: description,
		Frequency:   frequency,
		StartDate:   startDate,
		TargetDate:  targetDate,
		Priority:    priority,
	}
	return habit.ID, nil
}

func TestCreateHabit(t *testing.T) {

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		HabitService: *services.NewHabitService(&MockHabitStore{}),
		Sessions:     sessionManager,
	}

}
