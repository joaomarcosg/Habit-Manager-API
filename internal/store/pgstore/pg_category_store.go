package pgstore

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

var _ domain.CategoryRepository = (*PGCategoryStore)(nil)

func (pgc *PGCategoryStore) CreateCategory(ctx context.Context, category domain.Category) (uuid.UUID, error) {

	id, err := pgc.Queries.CreateCategory(ctx, category.Name)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, domain.ErrDuplicateCategoryName
		}
		return uuid.UUID{}, err
	}

	return id, nil
}

func (pgc *PGCategoryStore) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {

	category, err := pgc.Queries.GetCategoryByName(ctx, name)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Category{}, domain.ErrCategoryNotFound
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

func (pgc *PGCategoryStore) DeleteCategory(ctx context.Context, name string) (bool, error) {

	ok, err := pgc.Queries.DeleteCategory(ctx, name)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return false, domain.ErrCategoryInUse
		}
		return false, err
	}

	return ok.RowsAffected() > 0, nil

}
