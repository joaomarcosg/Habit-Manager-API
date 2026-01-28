package pgstore

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGHabitStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGHabitStore(pool *pgxpool.Pool) PGHabitStore {
	return PGHabitStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func (pgh *PGHabitStore) CreateHabit(
	ctx context.Context,
	name string,
	category Category,
	description string,
	frequency []Weekday,
	startDate,
	targetDate time.Time,
	priority int,
) (uuid.UUID, error) {
	id, err := pgh.Queries.CreateHabit(ctx, CreateHabitParams{
		Name:        name,
		Category:    category.Name,
		Description: description,
		Frequency:   frequency,
		StartDate:   startDate,
		TargetDate:  targetDate,
		Priority:    int16(priority),
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
