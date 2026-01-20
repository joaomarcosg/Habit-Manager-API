package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
)

var (
	ErrDuplicateCategoryName = errors.New("category name already exists")
	ErrCategoryNotFound      = errors.New("category not found")
)

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
			return uuid.UUID{}, ErrDuplicateCategoryName
		}
	}

	return id, nil
}

func (cs *CategoryService) GetCategoryById(ctx context.Context, id uuid.UUID) (store.Category, error) {
	category, err := cs.Store.GetCategoryById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return store.Category{}, ErrCategoryNotFound
		}
		return store.Category{}, err
	}

	return store.Category{
		ID:        category.ID,
		Name:      category.Name,
		Entries:   category.Entries,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}
