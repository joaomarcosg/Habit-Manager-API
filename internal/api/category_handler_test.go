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

type MockCategoryRepository struct {
	CreateCategoryFn    func(ctx context.Context, category domain.Category) (uuid.UUID, error)
	GetCategoryByNameFn func(ctx context.Context, name string) (domain.Category, error)
	DeleteCategoryFn    func(ctx context.Context, name string) (bool, error)
}

func (m *MockCategoryRepository) CreateCategory(ctx context.Context, category domain.Category) (uuid.UUID, error) {
	return m.CreateCategoryFn(ctx, category)
}

func (m *MockCategoryRepository) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {
	return m.GetCategoryByNameFn(ctx, name)
}

func (m *MockCategoryRepository) DeleteCategory(ctx context.Context, name string) (bool, error) {
	return m.DeleteCategoryFn(ctx, name)
}
func TestCreateCategory(t *testing.T) {

	mockRepo := &MockCategoryRepository{
		CreateCategoryFn: func(ctx context.Context, category domain.Category) (uuid.UUID, error) {
			return uuid.New(), nil
		},
	}

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		CategoryService: *services.NewCategoryService(mockRepo),
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

	mockRepo := &MockCategoryRepository{
		GetCategoryByNameFn: func(ctx context.Context, name string) (domain.Category, error) {
			return domain.Category{
				ID:        uuid.New(),
				Name:      "Health",
				Entries:   1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil
		},
	}

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		CategoryService: *services.NewCategoryService(mockRepo),
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

	mockRepo := &MockCategoryRepository{
		GetCategoryByNameFn: func(ctx context.Context, name string) (domain.Category, error) {
			return domain.Category{
				Entries: 1,
			}, nil
		},
	}

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		CategoryService: *services.NewCategoryService(mockRepo),
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

	mockRepo := &MockCategoryRepository{
		GetCategoryByNameFn: func(ctx context.Context, name string) (domain.Category, error) {
			return domain.Category{
				ID:      uuid.New(),
				Name:    "Health",
				Entries: 0,
			}, nil
		},
		DeleteCategoryFn: func(ctx context.Context, name string) (bool, error) {
			return true, nil
		},
	}

	gob.Register(uuid.UUID{})

	sessionManager := scs.New()
	sessionManager.Store = memstore.New()
	sessionManager.Lifetime = 1 * time.Hour

	api := Api{
		CategoryService: *services.NewCategoryService(mockRepo),
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
