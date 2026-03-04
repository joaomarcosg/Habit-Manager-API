package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	CreateUserFn       func(ctx context.Context, user domain.User) (uuid.UUID, error)
	GetUserByEmailFn   func(ctx context.Context, email string) (domain.User, error)
	GetUserByIdFn      func(ctx context.Context, id uuid.UUID) (domain.User, error)
	AuthenticateUserFn func(ctx context.Context, email, password string) (uuid.UUID, error)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user domain.User) (uuid.UUID, error) {
	return m.CreateUserFn(ctx, user)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return m.GetUserByEmailFn(ctx, email)
}

func (m *MockUserRepository) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return m.GetUserByIdFn(ctx, id)
}

func (m *MockUserRepository) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	return m.AuthenticateUserFn(ctx, email, password)
}

func TestSucessCreateUser(t *testing.T) {
	expectedID := uuid.New()

	mockRepo := &MockUserRepository{
		CreateUserFn: func(ctx context.Context, user domain.User) (uuid.UUID, error) {
			if user.Password == "123456" {
				t.Errorf("password should be hashed")
			}
			return expectedID, nil
		},
	}

	service := NewUserService(mockRepo)

	id, err := service.CreateUser(context.Background(), "John Doe", "johndoe@email.com", "Password123456")

	if err != nil {
		t.Fatalf("unexpected erro %v", err)
	}

	if id != expectedID {
		t.Fatalf("expected %v, got %v", expectedID, id)
	}
}

func TestDuplicateCreateUser(t *testing.T) {
	mockRepo := &MockUserRepository{
		CreateUserFn: func(ctx context.Context, user domain.User) (uuid.UUID, error) {
			return uuid.UUID{}, domain.ErrDuplicatedEmailOrUserName
		},
	}

	service := NewUserService(mockRepo)

	_, err := service.CreateUser(context.Background(), "John Doe", "johndoe@email.com", "Password123456")

	if !errors.Is(err, domain.ErrDuplicatedEmailOrUserName) {
		t.Fatalf("expected duplicate error, got %v", err)
	}
}

func TestSucessAuthenticateUser(t *testing.T) {
	password := "Password123456"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	expectedID := uuid.New()

	mockRepo := &MockUserRepository{
		GetUserByEmailFn: func(ctx context.Context, email string) (domain.User, error) {
			return domain.User{
				ID:       expectedID,
				Email:    email,
				Password: string(hash),
			}, nil
		},
	}

	service := NewUserService(mockRepo)

	id, err := service.AuthenticateUser(context.Background(), "johndoe@email.com", "Password123456")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id != expectedID {
		t.Fatalf("expected %v, got %v", expectedID, id)
	}
}

func TestInvalidPasswordAuthenticateUser(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("Password123456"), 12)

	mockRepo := &MockUserRepository{
		GetUserByEmailFn: func(ctx context.Context, email string) (domain.User, error) {
			return domain.User{
				ID:       uuid.New(),
				Email:    email,
				Password: string(hash),
			}, nil
		},
	}

	service := NewUserService(mockRepo)

	_, err := service.AuthenticateUser(context.Background(), "johndoe@email.com", "wrong-password123")

	if err == nil {
		t.Fatal("expected error, but got nil")
	}
}
