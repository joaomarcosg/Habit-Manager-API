package api

import (
	"context"

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
