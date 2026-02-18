package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
)

var (
	ErrDuplicateCategoryName = errors.New("category name already exists")
	ErrCategoryNotFound      = errors.New("category not found")
	ErrCategoryInUse         = errors.New("category is in use")
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

func (cs *CategoryService) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {
	category, err := cs.Store.GetCategoryByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Category{}, ErrCategoryNotFound
		}
		return domain.Category{}, err
	}

	return domain.Category{
		ID:        category.ID,
		Name:      category.Name,
		Entries:   category.Entries,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil

}

func (cs *CategoryService) GetCategoryEntries(ctx context.Context, name string) (domain.Category, error) {
	categoryEntries, err := cs.Store.GetCategoryByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Category{}, ErrCategoryNotFound
		}
		return domain.Category{}, err
	}

	return domain.Category{
		Entries: categoryEntries.Entries,
	}, nil
}

func (cs *CategoryService) DeleteCategory(ctx context.Context, name string) (bool, error) {
	category, err := cs.Store.GetCategoryByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, ErrCategoryNotFound
		}
		return false, err
	}

	ok, err := cs.Store.DeleteCategory(ctx, category.Name)
	var pgErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return false, ErrCategoryInUse
		}
	}

	return ok, nil

}
