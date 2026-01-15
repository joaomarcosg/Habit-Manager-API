package pgstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
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
