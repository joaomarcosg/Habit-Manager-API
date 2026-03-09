package domain

import (
	"context"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category Category) (uuid.UUID, error)
	GetCategoryByName(ctx context.Context, name string) (Category, error)
	DeleteCategory(ctx context.Context, name string) (bool, error)
}
