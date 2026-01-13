package pgstore

import "github.com/jackc/pgx/v5/pgxpool"

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
