package store

import "context"

type Category struct {
	Name    string `json:"category_name"`
	Entries int    `json:"entries"`
}

type CategoryStore interface {
	CreateCategory(ctx context.Context, name string) (Category, error)
	DeleteCategory(ctx context.Context, name string) (bool, error)
	GetCategoryByName(ctx context.Context, name string) (Category, error)
	GetCategoryEntries(ctx context.Context, name string) (Category, error)
}
