package pgstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

type PGCategoryStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGCategoryStore(pool *pgxpool.Pool) PGCategoryStore {
	return PGCategoryStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func (pgc *PGCategoryStore) CreateCategory(ctx context.Context, name string) (uuid.UUID, error) {
	id, err := pgc.Queries.CreateCategory(ctx, name)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (pgc *PGCategoryStore) GetCategoryById(ctx context.Context, id uuid.UUID) (domain.Category, error) {
	category, err := pgc.Queries.GetCategoryById(ctx, id)

	if err != nil {
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

func (pgc *PGCategoryStore) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {
	category, err := pgc.Queries.GetCategoryByName(ctx, name)

	if err != nil {
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

func (pgc *PGCategoryStore) GetCategoryEntries(ctx context.Context, name string) (domain.Category, error) {
	categoryEntries, err := pgc.Queries.GetCategoryEntries(ctx, name)

	if err != nil {
		return domain.Category{Entries: 0}, err
	}

	return domain.Category{Entries: categoryEntries}, nil
}

func (pgc *PGCategoryStore) DeleteCategory(ctx context.Context, id uuid.UUID) (bool, error) {
	ok, err := pgc.Queries.DeleteCategory(ctx, id)

	if err != nil {
		return false, err
	}

	return ok.Delete(), nil
}
