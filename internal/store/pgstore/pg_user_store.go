package pgstore

import (
	"context"

	"github.com/google/uuid"
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

func (pgu *PGUserStore) CreateUser(
	ctx context.Context,
	name,
	email string,
	password []byte,
) (uuid.UUID, error) {
	id, err := pgu.Queries.CreateUser(ctx, CreateUserParams{
		Name:         name,
		Email:        email,
		PasswordHash: password,
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (pgu *PGUserStore) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := pgu.Queries.GetUserByEmail(ctx, email)
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

func (pgu *PGUserStore) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	user, err := pgu.Queries.GetUserByEmail(ctx, email)
	if err != nil {
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
