package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
)

var ErrDuplicateCategoryName = errors.New("category name already exists")

type CategoryService struct {
	Store store.CategoryStore
}

func NewCategoryService(store store.CategoryStore) *CategoryService {
	return &CategoryService{
		Store: store,
	}
}

func (cs *CategoryService) CreateCategory(ctx context.Context, name string) (uuid.UUID, error) {
	id, err := cs.Store.CreateCategory(ctx, name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, err
		}
	}

	return id, nil
}
