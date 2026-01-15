package pgstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
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

func (pgc *PGCategoryStore) GetCategoryById(ctx context.Context, id uuid.UUID) (store.Category, error) {
	category, err := pgc.Queries.GetCategoryById(ctx, id)

	if err != nil {
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
