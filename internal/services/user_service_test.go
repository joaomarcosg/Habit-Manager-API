package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type MockUserStore struct{}

func (m *MockUserStore) CreateUser(
	ctx context.Context,
	name,
	email,
	password string,
) (uuid.UUID, error) {
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	return id, nil
}

func (m *MockUserStore) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	return id, nil
}

func (m *MockUserStore) GetUserByEmail(ctx context.Context, email string) (store.User, error) {
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")

	hash, _ := bcrypt.GenerateFromPassword([]byte("Senha123456"), bcrypt.DefaultCost)

	return store.User{
		ID:       id,
		Name:     "Fulano",
		Email:    "fulano@email.com",
		Password: string(hash),
	}, nil
}

func (m *MockUserStore) GetUserById(ctx context.Context, id uuid.UUID) (store.User, error) {
	return store.User{}, nil
}
