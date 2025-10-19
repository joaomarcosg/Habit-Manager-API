package service

import (
	"context"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicatedEmailOrUserName = errors.New("username or email already exists")
	ErrInvalidCredentials        = errors.New("invalid credentials")
)

type UserService struct {
	Store store.UserStore
}

func NewUserService(store store.UserStore) *UserService {
	return &UserService{
		Store: store,
	}
}

func (us *UserService) CreateUser(
	ctx context.Context,
	userName,
	email,
	password string,
) (uuid.UUID, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := us.Store.CreateUser(ctx, userName, email, passwordHash)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return uuid.UUID{}, ErrDuplicatedEmailOrUserName
		}
	}

	return id, nil
}
