package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) CreateUser(
	ctx context.Context,
	name,
	email,
	password string,
) (uuid.UUID, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.UUID{}, err
	}

	user := domain.User{
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	id, err := us.repo.CreateUser(ctx, user)

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (us *UserService) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {

	user, err := us.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return uuid.UUID{}, err
		}

		return uuid.UUID{}, err
	}

	return user.ID, nil

}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {

	user, err := us.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}

func (us *UserService) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {

	user, err := us.repo.GetUserById(ctx, id)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}
