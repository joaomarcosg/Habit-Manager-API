package services

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type MockUserStore struct{}

func (m *MockUserStore) CreateUser(
	ctx context.Context,
	name,
	email string,
	password []byte,
) (uuid.UUID, error) {
	id := uuid.New()
	return id, nil
}

func (m *MockUserStore) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	id := uuid.New()
	return id, nil
}

func (m *MockUserStore) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	id := uuid.New()

	hash, _ := bcrypt.GenerateFromPassword([]byte("Password123456"), bcrypt.DefaultCost)

	return domain.User{
		ID:       id,
		Name:     "John Doe",
		Email:    "johndoe@email.com",
		Password: string(hash),
	}, nil
}

func (m *MockUserStore) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return domain.User{
		ID:        id,
		Name:      "John Doe",
		Email:     "johndoe@email.com",
		Createdat: time.Now(),
		Updatedat: time.Now(),
	}, nil
}

func TestCreateUser(t *testing.T) {
	mockStore := MockUserStore{}
	userService := NewUserService(&mockStore)

	ctx := context.Background()
	userName := "John Doe"
	email := "johndoe@email.com"
	password, _ := bcrypt.GenerateFromPassword([]byte("Password123456"), bcrypt.DefaultCost)

	id, err := userService.Store.CreateUser(
		ctx,
		userName,
		email,
		password,
	)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, id)

}

func TestGetUserByEmail(t *testing.T) {
	mockStore := MockUserStore{}
	userService := NewUserService(&mockStore)

	ctx := context.Background()
	email := "johndoe@email.com"
	emptyUser := domain.User{}

	user, err := userService.Store.GetUserByEmail(ctx, email)

	assert.NoError(t, err)
	assert.NotEqual(t, emptyUser, user)
}

func TestGetUserById(t *testing.T) {
	mockStore := MockUserStore{}
	userService := NewUserService(&mockStore)

	ctx := context.Background()
	id := uuid.New()

	user, err := userService.Store.GetUserById(ctx, id)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, user.ID)

}
