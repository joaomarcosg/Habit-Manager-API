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

type MockCategoryStore struct{}

func (m *MockCategoryStore) CreateCategory(ctx context.Context, name string) (uuid.UUID, error) {
	category := domain.Category{
		ID:   uuid.New(),
		Name: name,
	}

	return category.ID, nil
}

func (m *MockCategoryStore) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {
	return domain.Category{}, nil
}

func (m *MockCategoryStore) GetCategoryEntries(ctx context.Context, name string) (int, error) {
	entries := 1
	return entries, nil
}

func (m *MockCategoryStore) DeleteCategory(ctx context.Context, name string) (bool, error) {
	return true, nil
}

func TestCreateCategory(t *testing.T) {

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		CategoryService: *services.NewCategoryService(&MockCategoryStore{}),
		Sessions:        sessionManager,
	}

	payLoad := map[string]any{
		"category_name": "Health",
	}

	body, err := json.Marshal(payLoad)
	if err != nil {
		t.Fatal("fail to parse request payload")
	}

	req := httptest.NewRequest("POST", "/api/v1/categories/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ctx, _ := sessionManager.Load(req.Context(), "")

	userID := uuid.New()
	sessionManager.Put(ctx, "AuthenticateUserId", userID)

	req = req.WithContext(ctx)

	handler := sessionManager.LoadAndSave(http.HandlerFunc(api.handleCreateCategory))
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusCreated {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusCreated)
	}
}

func TestGetCategoryByName(t *testing.T) {

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		CategoryService: *services.NewCategoryService(&MockCategoryStore{}),
		Sessions:        sessionManager,
	}

	payLoad := `{
		"category_name": "Health"
	}`

	body := []byte(payLoad)

	req := httptest.NewRequest("POST", "/api/v1/categories/getCategory", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ctx, _ := sessionManager.Load(req.Context(), "")

	userID := uuid.New()
	sessionManager.Put(ctx, "AuthenticateUserId", userID)

	req = req.WithContext(ctx)

	handler := sessionManager.LoadAndSave(http.HandlerFunc(api.handleGetCategoryByName))
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusOK {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusOK)
	}
}

func TestGetCategoryEntries(t *testing.T) {

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		CategoryService: *services.NewCategoryService(&MockCategoryStore{}),
		Sessions:        sessionManager,
	}

	payLoad := `{
		"category_name": "Health"
	}`

	body := []byte(payLoad)

	req := httptest.NewRequest("POST", "/api/v1/categories/getCategoryEntries", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ctx, _ := sessionManager.Load(req.Context(), "")

	userID := uuid.New()
	sessionManager.Put(ctx, "AuthenticateUserId", userID)

	req = req.WithContext(ctx)

	handler := sessionManager.LoadAndSave(http.HandlerFunc(api.handleGetCategoryEntries))
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusOK {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusOK)
	}

}

func TestDeleteCategory(t *testing.T) {

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		CategoryService: *services.NewCategoryService(&MockCategoryStore{}),
		Sessions:        sessionManager,
	}

	payLoad := `{
		"category_name": "Health"
	}`

	body := []byte(payLoad)

	req := httptest.NewRequest("POST", "/api/v1/categories/deleteCategory", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ctx, _ := sessionManager.Load(req.Context(), "")

	userID := uuid.New()
	sessionManager.Put(ctx, "AuthenticateUserId", userID)

	req = req.WithContext(ctx)

	handler := sessionManager.LoadAndSave(http.HandlerFunc(api.handleDeleteCategory))
	handler.ServeHTTP(rec, req)

	t.Logf("Rec body %s\n", rec.Body.Bytes())

	if rec.Code != http.StatusOK {
		t.Errorf("Statuscode differs; got %d | want %d", rec.Code, http.StatusOK)
	}

}
