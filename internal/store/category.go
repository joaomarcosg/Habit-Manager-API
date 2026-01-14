package store

import (
	"context"

	"github.com/google/uuid"
)

type Category struct {
	ID      uuid.UUID `json:"category_id"`
	Name    string    `json:"category_name"`
	Entries int       `json:"entries"`
}

type CategoryStore interface {
	CreateCategory(ctx context.Context, name string) (uuid.UUID, error)
	GetCategoryById(ctx context.Context, id uuid.UUID) (Category, error)
	GetCategoryByName(ctx context.Context, name string) (Category, error)
	GetCategoryEntries(ctx context.Context, name string) (Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (bool, error)
}
