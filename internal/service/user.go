package service

import (
	"context"

	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/repository"
)

type (
	UserService struct {
		userRepo *repository.UserRepository
	}
)

func NewUserService(
	userRepo *repository.UserRepository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.userRepo.CreateUser(ctx, user)
}

func (u *UserService) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	return u.userRepo.GetUserByID(ctx, userID)
}
