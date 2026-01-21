package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
)

type MockCategoryStore struct{}

func (m *MockCategoryStore) CreateCategory(ctx context.Context, name string) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (m *MockCategoryStore) GetCategoryById(ctx context.Context, id uuid.UUID) (store.Category, error) {
	return store.Category{}, nil
}

func (m *MockCategoryStore) GetCategoryByName(ctx context.Context, name string) (store.Category, error) {
	return store.Category{}, nil
}

func (m *MockCategoryStore) DeleteCategory(ctx context.Context, id uuid.UUID) (bool, error) {
	return true, nil
}

func TestCreateCategory(t *testing.T) {
	api := Api{}

	payLoad := map[string]any{
		"name": "Health",
	}

	body, err := json.Marshal(payLoad)
	if err != nil {
		t.Fatal("fail to parse request payload")
	}

	req := httptest.NewRequest("POST", "api/v1/categories", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(api.handleCreateCategory)
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusCreated {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusCreated)
	}
}
