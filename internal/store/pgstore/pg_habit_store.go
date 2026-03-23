package pgstore

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
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

var _ domain.HabitRepository = (*PGHabitStore)(nil)

func toPgTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}

func toPgWeekDays(weekDays []domain.WeekDay) []Weekday {
	days := make([]Weekday, 0, len(weekDays))

	for _, d := range weekDays {
		days = append(days, Weekday(d))
	}

	return days
}

func toDomainWeekDays(dbDays []Weekday) []domain.WeekDay {
	days := make([]domain.WeekDay, 0, len(dbDays))

	for _, d := range dbDays {
		days = append(days, domain.WeekDay(d))
	}

	return days
}

func toPgText(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}

	return pgtype.Text{
		String: s,
		Valid:  true,
	}
}

func toPgInt(i int) pgtype.Int2 {
	if i == 0 {
		return pgtype.Int2{Valid: false}
	}

	return pgtype.Int2{
		Int16: int16(i),
		Valid: true,
	}
}

func (pgh *PGHabitStore) CreateHabit(ctx context.Context, habit domain.Habit) (uuid.UUID, error) {

	tx, err := pgh.Pool.Begin(ctx)

	if err != nil {
		return uuid.UUID{}, err
	}

	defer tx.Rollback(ctx)

	qtx := pgh.Queries.WithTx(tx)

	id, err := qtx.CreateHabit(ctx, CreateHabitParams{
		Name:        habit.Name,
		Category:    habit.Category,
		Description: habit.Description,
		Frequency:   toPgWeekDays(habit.Frequency),
		StartDate:   toPgTimestamptz(habit.StartDate),
		TargetDate:  toPgTimestamptz(habit.TargetDate),
		Priority:    int16(habit.Priority),
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	tag, err := qtx.IncrementCategoryEntries(ctx, habit.Category)
	if err != nil {
		return uuid.UUID{}, err
	}

	if tag.RowsAffected() == 0 {
		return uuid.Nil, domain.ErrCategoryNotFound
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.UUID{}, err
	}

	return id, nil

}

func (pgh *PGHabitStore) GetHabitByName(ctx context.Context, name string) (domain.Habit, error) {

	habit, err := pgh.Queries.GetHabitByName(ctx, name)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Habit{}, domain.ErrHabitNotFound
		}
		return domain.Habit{}, err
	}

	return domain.Habit{
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

func (pgh *PGHabitStore) UpdateHabit(ctx context.Context, habit domain.Habit) (domain.Habit, error) {

	updatedHabit, err := pgh.Queries.UpdateHabit(ctx, UpdateHabitParams{
		Name:        toPgText(habit.Name),
		Category:    toPgText(habit.Category),
		Description: toPgText(habit.Description),
		Frequency:   toPgWeekDays(habit.Frequency),
		StartDate:   toPgTimestamptz(habit.StartDate),
		TargetDate:  toPgTimestamptz(habit.TargetDate),
		Priority:    toPgInt(habit.Priority),
	})

	if err != nil {
		return domain.Habit{}, err
	}

	return domain.Habit{
		ID:          updatedHabit.ID,
		Name:        updatedHabit.Name,
		Category:    updatedHabit.Category,
		Description: updatedHabit.Description,
		Frequency:   toDomainWeekDays(updatedHabit.Frequency),
		StartDate:   updatedHabit.StartDate.Time,
		TargetDate:  updatedHabit.TargetDate.Time,
		Priority:    int(updatedHabit.Priority),
		CreatedAt:   updatedHabit.CreatedAt,
		UpdatedAt:   updatedHabit.UpdatedAt,
	}, nil
}

func (pgh *PGHabitStore) DeleteHabit(ctx context.Context, name string) error {

	_, err := pgh.Queries.DeleteHabit(ctx, name)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrHabitNotFound
		}

		return err
	}

	return nil
}
