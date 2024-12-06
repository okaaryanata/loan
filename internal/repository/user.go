package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/helper"
	"github.com/okaaryanata/loan/internal/repository/query"
)

type (
	UserRepository struct {
		db *pgxpool.Pool
	}
)

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := query.QueryCreateUser
	err := u.db.QueryRow(ctx, query, user.Username, user.IsActive, user.OperatedBy).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	row := u.db.QueryRow(ctx, query.QueryGetUserByID, userID)
	user := &domain.User{}
	err := helper.StructScan(row, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	row := u.db.QueryRow(ctx, query.QueryGetUserByUsername, username)
	user := &domain.User{}
	err := helper.StructScan(row, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
