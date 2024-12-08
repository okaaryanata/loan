package service

import (
	"context"
	"net/http"

	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/helper"
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

func (u *UserService) CreateUser(ctx context.Context, args *domain.UserRequest) (*domain.User, helper.Errorx) {
	newUser := &domain.User{
		Username:   args.Username,
		OperatedBy: helper.Chains(args.OperatedBy, "SYSTEM"),
		IsActive:   true,
	}

	existingUser, err := u.userRepo.GetUserByUsernames(ctx, args.Username)
	if err == nil && len(existingUser) > 0 {
		return nil, helper.NewErrorxif(http.StatusBadRequest, "user with username %s already exists", args.Username)
	}

	err = u.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	return newUser, nil
}

func (u *UserService) GetUserByID(ctx context.Context, userID int64) (*domain.User, helper.Errorx) {
	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	if user == nil {
		return nil, helper.NewErrorx(http.StatusNotFound, "user not found")
	}

	return user, nil
}

func (u *UserService) GetUserByUsernames(ctx context.Context, username ...string) ([]*domain.User, helper.Errorx) {
	users, err := u.userRepo.GetUserByUsernames(ctx, username...)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	return users, nil
}

func (u *UserService) GetUsers(ctx context.Context, username ...string) ([]*domain.User, helper.Errorx) {
	if len(username) > 0 {
		users, err := u.GetUserByUsernames(ctx, username...)
		if err != nil {
			return nil, err
		}

		return users, nil
	}

	users, err := u.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	return users, nil
}
