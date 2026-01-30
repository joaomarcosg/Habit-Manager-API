package pgstore

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
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

func toPgTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}

func toDomainWeekDays(dbDays []Weekday) []store.WeekDay {
	days := make([]store.WeekDay, 0, len(dbDays))

	for _, d := range dbDays {
		days = append(days, store.WeekDay(d))
	}

	return days
}

func (pgh *PGHabitStore) CreateHabit(
	ctx context.Context,
	name string,
	category,
	description string,
	frequency []Weekday,
	startDate,
	targetDate time.Time,
	priority int,
) (uuid.UUID, error) {
	id, err := pgh.Queries.CreateHabit(ctx, CreateHabitParams{
		Name:        name,
		Category:    category,
		Description: description,
		Frequency:   frequency,
		StartDate:   toPgTimestamptz(startDate),
		TargetDate:  toPgTimestamptz(targetDate),
		Priority:    int16(priority),
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (pgh *PGHabitStore) GetHabitById(ctx context.Context, id uuid.UUID) (store.Habit, error) {
	habit, err := pgh.Queries.GetHabitById(ctx, id)

	if err != nil {
		return store.Habit{}, err
	}

	return store.Habit{
		ID:          habit.ID,
		Name:        habit.Name,
		Category:    habit.Category,
		Description: habit.Description,
		Frequency:   toDomainWeekDays(habit.Frequency),
		StartDate:   habit.StartDate.Time,
		TargetDate:  habit.TargetDate.Time,
		Priority:    int(habit.Priority),
		CreatedAt:   habit.CreatedAt,
		UpdatedAt:   habit.UpdatedAt,
	}, nil
}

func (pgh *PGHabitStore) GetHabitByName(ctx context.Context, name string) (store.Habit, error) {
	habit, err := pgh.Queries.GetHabitByName(ctx, name)

	if err != nil {
		return store.Habit{}, err
	}

	return store.Habit{
		ID:          habit.ID,
		Name:        habit.Name,
		Category:    habit.Category,
		Description: habit.Description,
		Frequency:   toDomainWeekDays(habit.Frequency),
		StartDate:   habit.StartDate.Time,
		TargetDate:  habit.TargetDate.Time,
		Priority:    int(habit.Priority),
		CreatedAt:   habit.CreatedAt,
		UpdatedAt:   habit.UpdatedAt,
	}, nil
}

func (pgh *PGHabitStore) UpdateHabit(
	ctx context.Context,
	name,
	category,
	description string,
	frequency []Weekday,
	startDate,
	targetDate time.Time,
	priority int,
) (store.Habit, error) {
	updatedHabit, err := pgh.Queries.UpdateHabit(ctx, UpdateHabitParams{
		Name:        name,
		Category:    category,
		Description: description,
		Frequency:   frequency,
		StartDate:   startDate,
		TargetDate:  targetDate,
		Priority:    priority,
	})

	if err != nil {
		return store.Habit{}, err
	}

	return store.Habit{}, nil
}
