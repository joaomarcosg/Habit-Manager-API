package mysqlstore

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
)

type MSUserStore struct {
	Queries *Queries
	Pool    *sql.DB
}

func NewMSUserStore(pool *sql.DB) MSUserStore {
	return MSUserStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func (msu *MSUserStore) CreateUser(
	ctx context.Context,
	userName,
	email,
	password string,
) (uuid.UUID, error) {
	_, err := msu.Queries.CreateUser(ctx, CreateUserParams{
		Name:     userName,
		Email:    email,
		Password: password,
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	id := uuid.New()

	return id, nil
}

func (msu *MSUserStore) GetUserByEmail(ctx context.Context, email string) (store.User, error) {
	user, err := msu.Queries.GetUserByEmail(ctx, email)

	if err != nil {
		return store.User{}, err
	}

	return store.User{
		ID:        uuid.UUID(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Createdat: user.CreatedAt.Time,
		Updatedat: user.UpdatedAt.Time,
	}, nil
}

func (msu *MSUserStore) GetUserById(ctx context.Context, id uuid.UUID) (store.User, error) {
	user, err := msu.Queries.GetUserById(ctx, id[:])

	if err != nil {
		return store.User{}, err
	}

	return store.User{
		ID:        uuid.UUID(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Createdat: user.CreatedAt.Time,
		Updatedat: user.UpdatedAt.Time,
	}, nil
}
