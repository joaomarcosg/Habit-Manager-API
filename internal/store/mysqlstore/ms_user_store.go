package mysqlstore

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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
