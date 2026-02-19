package services

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockCategoryStore struct{}

func (m *MockCategoryStore) CreateCategory(ctx context.Context, name string) (uuid.UUID, error) {
	id := uuid.New()
	return id, nil
}

func (m *MockCategoryStore) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {
	id := uuid.New()
	return domain.Category{
		ID:        id,
		Name:      name,
		Entries:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *MockCategoryStore) GetCategoryEntries(ctx context.Context, name string) (int, error) {
	entries := 1
	return entries, nil
}

func (m *MockCategoryStore) DeleteCategory(ctx context.Context, name string) (bool, error) {
	return true, nil
}

func TestCreateCategory(t *testing.T) {
	mockStore := MockCategoryStore{}
	categoryService := NewCategoryService(&mockStore)

	id, err := categoryService.Store.CreateCategory(context.Background(), "Health")

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

}

func TestGetCategoryByName(t *testing.T) {
	mockStore := MockCategoryStore{}
	categoryService := NewCategoryService(&mockStore)

	ctx := context.Background()
	name := "Health"
	emptyCategory := domain.Category{}

	category, err := categoryService.Store.GetCategoryByName(ctx, name)

	assert.NoError(t, err)
	assert.NotEqual(t, emptyCategory, category)
	assert.Equal(t, name, category.Name)
}

func TestGetCategoryEntries(t *testing.T) {
	mockStore := MockCategoryStore{}
	categoryService := NewCategoryService(&mockStore)

	ctx := context.Background()
	name := "Health"
	emptyCategory := domain.Category{}

	categoryEntries, err := categoryService.GetCategoryByName(ctx, name)

	assert.NoError(t, err)
	assert.NotEqual(t, emptyCategory, categoryEntries)
	assert.Equal(t, 1, categoryEntries.Entries)

}

func TestDeleteCategory(t *testing.T) {
	mockStore := MockCategoryStore{}
	categoryService := NewCategoryService(&mockStore)

	ctx := context.Background()
	name := "Health"

	category, err := categoryService.Store.GetCategoryByName(ctx, name)
	if category.Name != name {
		t.Errorf("%s not found", category.ID)
	}

	ok, err := categoryService.Store.DeleteCategory(ctx, name)

	assert.NoError(t, err)
	assert.Equal(t, ok, true)

}
