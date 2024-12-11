package domain

import "time"

type (
	User struct {
		ID         int64      `json:"userID" db:"user_id"`
		Username   string     `json:"username" db:"username"`
		IsActive   bool       `json:"isActive" db:"is_active"`
		CreatedAt  time.Time  `json:"-" db:"created_at"`
		CreatedBy  string     `json:"-" db:"created_by"`
		UpdatedAt  time.Time  `json:"-" db:"updated_at"`
		UpdatedBy  string     `json:"-" db:"updated_by"`
		DeletedAt  *time.Time `json:"-" db:"deleted_at"`
		DeletedBy  *string    `json:"-" db:"deleted_by"`
		OperatedBy string     `json:"-" db:"-"`
	}

	UserRequest struct {
		Username   string `json:"username"`
		OperatedBy string `json:"operateBy"`
	}
)
