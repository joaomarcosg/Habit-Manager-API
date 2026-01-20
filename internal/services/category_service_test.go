package services

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
	"github.com/stretchr/testify/assert"
)

type MockCategoryStore struct{}

func (m *MockCategoryStore) CreateCategory(ctx context.Context, name string) (uuid.UUID, error) {
	id := uuid.New()
	return id, nil
}

func (m *MockCategoryStore) GetCategoryById(ctx context.Context, id uuid.UUID) (store.Category, error) {
	return store.Category{
		ID:        id,
		Name:      "Health",
		Entries:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *MockCategoryStore) GetCategoryByName(ctx context.Context, name string) (store.Category, error) {
	id := uuid.New()
	return store.Category{
		ID:        id,
		Name:      name,
		Entries:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *MockCategoryStore) GetCategoryEntries(ctx context.Context, name string) (store.Category, error) {
	return store.Category{
		Entries: 1,
	}, nil
}

func (m *MockCategoryStore) DeleteCategory(ctx context.Context, id uuid.UUID) (bool, error) {
	return true, nil
}

func TestCreateCategory(t *testing.T) {
	mockStore := MockCategoryStore{}
	categoryService := NewCategoryService(&mockStore)

	id, err := categoryService.Store.CreateCategory(context.Background(), "Health")

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

}

func TestGetCategoryById(t *testing.T) {
	mockStore := MockCategoryStore{}
	categoryService := NewCategoryService(&mockStore)

	ctx := context.Background()
	id := uuid.New()
	emptyCategory := store.Category{}

	category, err := categoryService.Store.GetCategoryById(ctx, id)

	assert.NoError(t, err)
	assert.Equal(t, id, category.ID)
	assert.NotEqual(t, emptyCategory, category)

}

func TestGetCategoryByName(t *testing.T) {
	mockStore := MockCategoryStore{}
	categoryService := NewCategoryService(&mockStore)

	ctx := context.Background()
	name := "Health"
	emptyCategory := store.Category{}

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
	emptyCategory := store.Category{}

	categoryEntries, err := categoryService.GetCategoryByName(ctx, name)

	assert.NoError(t, err)
	assert.NotEqual(t, emptyCategory, categoryEntries)
	assert.Equal(t, 1, categoryEntries.Entries)

}

func TestDeleteCategory(t *testing.T) {
	mockStore := MockCategoryStore{}
	categoryService := NewCategoryService(&mockStore)

	ctx := context.Background()
	id := uuid.New()

	category, err := categoryService.Store.GetCategoryById(ctx, id)
	if category.ID != id {
		t.Errorf("%s not found", category.ID)
	}

	ok, err := categoryService.Store.DeleteCategory(ctx, id)

	assert.NoError(t, err)
	assert.Equal(t, ok, true)

}
