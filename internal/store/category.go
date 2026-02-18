package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

type CategoryStore interface {
	CreateCategory(ctx context.Context, name string) (uuid.UUID, error)
	GetCategoryByName(ctx context.Context, name string) (domain.Category, error)
	GetCategoryEntries(ctx context.Context, name string) (int, error)
	DeleteCategory(ctx context.Context, name string) (bool, error)
}
