package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrDuplicateCategoryName = errors.New("category name already exists")
	ErrCategoryNotFound      = errors.New("category not found")
	ErrCategoryInUse         = errors.New("category is in use")
)

type Category struct {
	ID        uuid.UUID
	Name      string
	Entries   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
