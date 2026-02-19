package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
)

var (
	ErrDuplicateHabitName = errors.New("habit name already exists")
	ErrHabitNotFound      = errors.New("habit not found")
)

type HabitService struct {
	Store store.HabitStore
}

func NewHabitService(store store.HabitStore) *HabitService {
	return &HabitService{
		Store: store,
	}
}

func (hs *HabitService) CreateHabit(
	ctx context.Context,
	name,
	category,
	description string,
	frequency []domain.WeekDay,
	startDate,
	targetDate time.Time,
	priority int,
) (uuid.UUID, error) {

	id, err := hs.Store.CreateHabit(
		ctx,
		name,
		category,
		description,
		frequency,
		startDate,
		targetDate,
		priority,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicateHabitName
		}

		return uuid.UUID{}, err
	}

	return id, nil

}

func (hs *HabitService) GetHabitById(ctx context.Context, id uuid.UUID) (domain.Habit, error) {

	habit, err := hs.Store.GetHabitById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Habit{}, ErrHabitNotFound
		}
		return domain.Habit{}, err
	}

	return domain.Habit{
		ID:          habit.ID,
		Name:        habit.Name,
		Category:    habit.Category,
		Description: habit.Description,
		Frequency:   habit.Frequency,
		StartDate:   habit.StartDate,
		TargetDate:  habit.TargetDate,
		Priority:    habit.Priority,
	}, nil

}
