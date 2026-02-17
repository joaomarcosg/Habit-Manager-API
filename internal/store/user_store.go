package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
)

type UserStore interface {
	CreateUser(ctx context.Context, userName, email string, password []byte) (uuid.UUID, error)
	AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error)
}
