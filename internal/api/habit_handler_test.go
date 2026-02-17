package api

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
	"github.com/joaomarcosg/Habit-Manager-API/internal/services"
)

type MockHabitStore struct{}

// DeleteHabit implements store.HabitStore.
func (m *MockHabitStore) DeleteHabit(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	panic("unimplemented")
}

// GetHabitById implements store.HabitStore.
func (m *MockHabitStore) GetHabitById(ctx context.Context, id uuid.UUID) (domain.Habit, error) {
	panic("unimplemented")
}

// GetHabitByName implements store.HabitStore.
func (m *MockHabitStore) GetHabitByName(ctx context.Context, name string) (domain.Habit, error) {
	panic("unimplemented")
}

// UpdateHabit implements store.HabitStore.
func (m *MockHabitStore) UpdateHabit(
	ctx context.Context,
	name,
	category,
	description string,
	frequency []domain.WeekDay,
	startDate,
	targetDate time.Time,
) (domain.Habit, error) {
	panic("unimplemented")
}

func (m *MockHabitStore) CreateHabit(
	ctx context.Context,
	name,
	category,
	description string,
	frequency []domain.WeekDay,
	startDate,
	targetDate time.Time,
	priority int,
) (uuid.UUID, error) {
	habit := domain.Habit{
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

	payLoad := `{
		"habit_name": "Work out",
		"habit_category": "Health",
		"description": "Work out 5 times a week",
		"frequency": ["monday", "tuesday", "wednesday", "thursday", "friday"],
		"start_date": "2026-02-13T00:00:00Z",
		"target_date": "2026-02-20T00:00:00Z",
		"priority": 10,
	}`

	body, err := json.Marshal(payLoad)
	if err != nil {
		t.Fatal("fail to parse request payload")
	}

	req := httptest.NewRequest("POST", "/api/v1/habits/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ctx, _ := sessionManager.Load(req.Context(), "")

	userID := uuid.New()
	sessionManager.Put(ctx, "AuthenticateUserId", userID)

	req = req.WithContext(ctx)

	handler := sessionManager.LoadAndSave(http.HandlerFunc(api.handleCreateHabit))
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusCreated {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusCreated)
	}

}
