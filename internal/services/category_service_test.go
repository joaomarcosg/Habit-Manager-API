package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
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

func TestSuccessCreateCategory(t *testing.T) {
	expectedID := uuid.New()

	mockRepo := &MockCategoryRepository{
		CreateCategoryFn: func(ctx context.Context, category domain.Category) (uuid.UUID, error) {
			return expectedID, nil
		},
	}

	service := NewCategoryService(mockRepo)

	id, err := service.CreateCategory(context.Background(), "Health")

	if err != nil {
		t.Fatalf("unexpected erro %v", err)
	}

	if id != expectedID {
		t.Fatalf("expected %v, got %v", expectedID, id)
	}
}

func TestDuplicateNameCreateCategory(t *testing.T) {
	mockRepo := &MockCategoryRepository{
		CreateCategoryFn: func(ctx context.Context, category domain.Category) (uuid.UUID, error) {
			return uuid.UUID{}, domain.ErrDuplicateCategoryName
		},
	}

	service := NewCategoryService(mockRepo)

	id, err := service.CreateCategory(context.Background(), "Health")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, domain.ErrDuplicateCategoryName) {
		t.Fatalf("expected ErrDuplicateCategoryName, got %v", err)
	}

	if id != uuid.Nil {
		t.Fatalf("expcted empty uuid, got %v", id)
	}
}

func TestSuccessGetCategoryByName(t *testing.T) {
	expectedCategory := domain.Category{
		ID:        uuid.New(),
		Name:      "Health",
		Entries:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo := &MockCategoryRepository{
		GetCategoryByNameFn: func(ctx context.Context, name string) (domain.Category, error) {
			return domain.Category{
				ID:        expectedCategory.ID,
				Name:      expectedCategory.Name,
				Entries:   expectedCategory.Entries,
				CreatedAt: expectedCategory.CreatedAt,
				UpdatedAt: expectedCategory.UpdatedAt,
			}, nil
		},
	}

	service := NewCategoryService(mockRepo)

	category, err := service.GetCategoryByName(context.Background(), "Health")

	if err != nil {
		t.Fatalf("unexpected erro %v", err)
	}

	if category != expectedCategory {
		t.Fatalf("expected %v, got %v", expectedCategory, category)
	}

}

func TestCategoryNotFoundGetCategoryByName(t *testing.T) {
	emptyCategory := domain.Category{}

	mockRepo := &MockCategoryRepository{
		GetCategoryByNameFn: func(ctx context.Context, name string) (domain.Category, error) {
			return domain.Category{}, domain.ErrCategoryNotFound
		},
	}

	service := NewCategoryService(mockRepo)

	category, err := service.GetCategoryByName(context.Background(), "Healht")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, domain.ErrCategoryNotFound) {
		t.Fatalf("expected ErrCategoryNotFound, got %v", err)
	}

	if category != emptyCategory {
		t.Fatalf("expected empty category, got %v", category)
	}

}

func TestGetCategoryEntries(t *testing.T) {
	expectedCategory := domain.Category{
		ID:        uuid.New(),
		Name:      "Health",
		Entries:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	expectedEntries := 1

	mockRepo := &MockCategoryRepository{
		GetCategoryByNameFn: func(ctx context.Context, name string) (domain.Category, error) {
			return domain.Category{
				ID:        expectedCategory.ID,
				Name:      expectedCategory.Name,
				Entries:   expectedCategory.Entries,
				CreatedAt: expectedCategory.CreatedAt,
				UpdatedAt: expectedCategory.UpdatedAt,
			}, nil
		},
	}

	service := NewCategoryService(mockRepo)

	categoryEntries, err := service.GetCategoryEntries(context.Background(), "Health")

	if err != nil {
		t.Fatalf("unexpected erro %v", err)
	}

	if categoryEntries.Entries != expectedEntries {
		t.Fatalf("expected %d, got %d", expectedEntries, categoryEntries.Entries)
	}

}

func TestSuccessDeleteCategory(t *testing.T) {
	deleteCalled := false

	mockRepo := &MockCategoryRepository{
		GetCategoryByNameFn: func(ctx context.Context, name string) (domain.Category, error) {
			return domain.Category{
				ID:      uuid.New(),
				Name:    "Health",
				Entries: 0,
			}, nil
		},
		DeleteCategoryFn: func(ctx context.Context, name string) (bool, error) {
			deleteCalled = true
			return true, nil
		},
	}

	service := NewCategoryService(mockRepo)

	ok, err := service.DeleteCategory(context.Background(), "Health")

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !ok {
		t.Fatalf("expcted true, got %v", err)
	}

	if !deleteCalled {
		t.Fatalf("DeleteCategory should have been called")
	}
}

func TestCategoryInUseDeleteCategory(t *testing.T) {

	mockRepo := MockCategoryRepository{
		DeleteCategoryFn: func(ctx context.Context, name string) (bool, error) {
			return true, domain.ErrCategoryInUse
		},
	}

	service := NewCategoryService(&mockRepo)

	ok, err := service.DeleteCategory(context.Background(), "Health")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, domain.ErrCategoryInUse) {
		t.Fatalf("expected ErrCategoryInUse, got %v", err)
	}

	if ok {
		t.Fatalf("expected false, got %v", ok)
	}

}
