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

type PGUserStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGUserStore(pool *pgxpool.Pool) PGUserStore {
	return PGUserStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

var _ domain.UserRepository = (*PGUserStore)(nil)

func (pgu *PGUserStore) CreateUser(ctx context.Context, user domain.User) (uuid.UUID, error) {

	id, err := pgu.Queries.CreateUser(ctx, CreateUserParams{
		user.Name,
		user.Email,
		[]byte(user.Password),
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, domain.ErrDuplicatedEmailOrUserName
		}
		return uuid.UUID{}, err
	}

	return id, nil
}

func (pgu *PGUserStore) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {

	user, err := pgu.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Createdat: user.CreatedAt,
		Updatedat: user.UpdatedAt,
	}, nil
}

func (pgu *PGUserStore) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {

	user, err := pgu.Queries.GetUserByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, domain.ErrInvalidCredentials
		}
		return uuid.UUID{}, err
	}

	return user.ID, nil
}

func (pgu *PGUserStore) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	user, err := pgu.Queries.GetUserById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Createdat: user.CreatedAt,
		Updatedat: user.UpdatedAt,
	}, nil
}
