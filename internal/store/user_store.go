package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
}

type UserStore interface {
	CreateUser(ctx context.Context, userName, email string, password []byte) (uuid.UUID, error)
	AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (User, error)
}
