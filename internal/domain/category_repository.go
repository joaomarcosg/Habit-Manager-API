package domain

import (
	"context"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, name string) (uuid.UUID, error)
	GetCategoryByName(ctx context.Context, name string) (Category, error)
	GetCategoryEntries(ctx context.Context, name string) (int, error)
	DeleteCategory(ctx context.Context, name string) (bool, error)
}
