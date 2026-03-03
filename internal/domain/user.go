package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrDuplicatedEmailOrUserName = errors.New("username or email already exists")
	ErrInvalidCredentials        = errors.New("invalid credentials")
	ErrUserNotFound              = errors.New("user not found")
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	Createdat time.Time
	Updatedat time.Time
}
