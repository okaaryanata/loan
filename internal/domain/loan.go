package domain

import "time"

type (
	Loan struct {
		ID                 int64     `json:"loanID" db:"loan_id"`
		UserID             int64     `json:"userID" db:"user_id"`
		Code               string    `json:"code" db:"code"`
		Principal          float64   `json:"principal" db:"principal"`
		InterestRate       float64   `json:"interestRate" db:"interest_rate"`
		TotalWeeks         int       `json:"totalWeeks" db:"total_weeks"`
		WeeklyRepayment    float64   `json:"weeklyRepayment" db:"weekly_repayment"`
		OutstandingBalance float64   `json:"outstandingBalance" db:"outstanding_balance"`
		MissedPayments     int       `json:"missedPayments" db:"missed_payments"`
		IsDelinquent       bool      `json:"isDelinquent" db:"is_delinquent"`
		IsActive           bool      `json:"isActive" db:"is_active"`
		CreatedAt          time.Time `json:"-" db:"created_at"`
		CreatedBy          string    `json:"-" db:"created_by"`
		UpdatedAt          time.Time `json:"-" db:"updated_at"`
		UpdatedBy          string    `json:"-" db:"updated_by"`
		DeletedAt          time.Time `json:"-" db:"deleted_at"`
		DeletedBy          string    `json:"-" db:"deleted_by"`
		OperatedBy         string    `json:"-" db:"-"`
	}

	LoanRepayment struct {
		ID         int64     `json:"loanRepaymentID" db:"loan_repayment_id"`
		LoanID     int64     `json:"loanID" db:"loan_id"`
		Week       int       `json:"week" db:"week"`
		Amount     float64   `json:"amount" db:"amount"`
		Paid       bool      `json:"paid" db:"paid"`
		DueDate    time.Time `json:"dueDate" db:"due_date"`
		IsActive   bool      `json:"isActive" db:"is_active"`
		CreatedAt  time.Time `json:"-" db:"created_at"`
		CreatedBy  string    `json:"-" db:"created_by"`
		UpdatedAt  time.Time `json:"-" db:"updated_at"`
		UpdatedBy  string    `json:"-" db:"updated_by"`
		DeletedAt  time.Time `json:"-" db:"deleted_at"`
		DeletedBy  string    `json:"-" db:"deleted_by"`
		OperatedBy string    `json:"-" db:"-"`
	}

	LoanRequest struct {
		UserID       int64   `json:"userID"`
		Code         string  `json:"code"`
		Principal    float64 `json:"principal"`
		InterestRate float64 `json:"interestRate"`
		TotalWeeks   int     `json:"totalWeeks"`
		OperatedBy   string  `json:"operateBy"`
	}

	MakePaymentRequest struct {
		UserID      int64   `json:"userID"`
		LoanID      int64   `json:"loanID"`
		RepaymentID int64   `json:"repaymentID"`
		Amount      float64 `json:"amount"`
		Paid        bool    `json:"-"`
		OperatedBy  string  `json:"operatedBy"`
	}
)
