package repository

import (
	"context"

	"github.com/okaaryanata/loan/internal/domain"
)

type (
	UserRepository struct{}
)

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (u *UserRepository) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	return nil, nil
}
