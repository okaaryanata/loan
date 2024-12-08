package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okaaryanata/loan/internal/domain"
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
	user := &domain.User{}
	err := u.db.QueryRow(ctx, query.QueryGetUserByID, userID).Scan(
		&user.ID,
		&user.Username,
		&user.IsActive,
		&user.CreatedBy,
		&user.CreatedAt,
		&user.UpdatedBy,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserByUsernames(ctx context.Context, username ...string) ([]*domain.User, error) {
	rows, err := u.db.Query(ctx, query.QueryGetUserByUsernames, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.IsActive,
			&user.CreatedBy,
			&user.CreatedAt,
			&user.UpdatedBy,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) GetUsers(ctx context.Context) ([]*domain.User, error) {
	rows, err := u.db.Query(ctx, query.QueryGetUsers)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.IsActive,
			&user.CreatedBy,
			&user.CreatedAt,
			&user.UpdatedBy,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
