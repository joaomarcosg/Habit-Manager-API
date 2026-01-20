package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"category_id"`
	Name      string    `json:"category_name"`
	Entries   int       `json:"entries"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryStore interface {
	CreateCategory(ctx context.Context, name string) (uuid.UUID, error)
	GetCategoryById(ctx context.Context, id uuid.UUID) (Category, error)
	GetCategoryByName(ctx context.Context, name string) (Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (bool, error)
}
