package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/helper"
	"github.com/okaaryanata/loan/internal/repository/query"
)

type (
	LoanRepository struct {
		db *pgxpool.Pool
	}
)

func NewLoanRepository(db *pgxpool.Pool) *LoanRepository {
	return &LoanRepository{
		db: db,
	}
}

func (l *LoanRepository) CreateLoan(ctx context.Context, tx pgx.Tx, loan *domain.Loan) error {
	var err error
	if tx != nil {
		err = tx.QueryRow(ctx, query.QueryCreateLoan,
			loan.UserID,
			loan.Code,
			loan.Principal,
			loan.InterestRate,
			loan.TotalWeeks,
			loan.WeeklyRepayment,
			loan.OutstandingBalance,
			loan.MissedPayments,
			loan.IsDelinquent,
			loan.IsActive,
			loan.OperatedBy,
		).Scan(&loan.ID)
	} else {
		err = l.db.QueryRow(ctx, query.QueryCreateLoan,
			loan.UserID,
			loan.Code,
			loan.Principal,
			loan.InterestRate,
			loan.TotalWeeks,
			loan.WeeklyRepayment,
			loan.OutstandingBalance,
			loan.MissedPayments,
			loan.IsDelinquent,
			loan.IsActive,
			loan.OperatedBy,
		).Scan(&loan.ID)
	}
	if err != nil {
		return err
	}

	return nil
}

func (l *LoanRepository) GetLoanByID(ctx context.Context, loanID int64) (*domain.Loan, error) {
	loan := &domain.Loan{}
	err := l.db.QueryRow(ctx, query.QueryGetLoanByID, loanID).Scan(
		&loan.ID,
		&loan.Code,
		&loan.UserID,
		&loan.Principal,
		&loan.InterestRate,
		&loan.TotalWeeks,
		&loan.WeeklyRepayment,
		&loan.OutstandingBalance,
		&loan.MissedPayments,
		&loan.IsDelinquent,
		&loan.IsActive,
		&loan.CreatedBy,
		&loan.CreatedAt,
		&loan.UpdatedBy,
		&loan.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return loan, nil
}

func (l *LoanRepository) GetLoansByUserID(ctx context.Context, userID int64) ([]*domain.Loan, error) {
	rows, err := l.db.Query(ctx, query.QueryGetLoansByUserID, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	loans := []*domain.Loan{}
	for rows.Next() {
		var loan domain.Loan
		err := rows.Scan(
			&loan.ID,
			&loan.Code,
			&loan.UserID,
			&loan.Principal,
			&loan.InterestRate,
			&loan.TotalWeeks,
			&loan.WeeklyRepayment,
			&loan.OutstandingBalance,
			&loan.MissedPayments,
			&loan.IsDelinquent,
			&loan.IsActive,
			&loan.CreatedBy,
			&loan.CreatedAt,
			&loan.UpdatedBy,
			&loan.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		loans = append(loans, &loan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return loans, nil
}

func (l *LoanRepository) GetLoansByIDandUserID(ctx context.Context, loanID, userID int64) (*domain.Loan, error) {
	loan := &domain.Loan{}
	err := l.db.QueryRow(ctx, query.QueryGetLoanByIDandUserID, loanID, userID).Scan(
		&loan.ID,
		&loan.Code,
		&loan.UserID,
		&loan.Principal,
		&loan.InterestRate,
		&loan.TotalWeeks,
		&loan.WeeklyRepayment,
		&loan.OutstandingBalance,
		&loan.MissedPayments,
		&loan.IsDelinquent,
		&loan.IsActive,
		&loan.CreatedBy,
		&loan.CreatedAt,
		&loan.UpdatedBy,
		&loan.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return loan, nil
}

func (l *LoanRepository) UpdateLoan(ctx context.Context, loan *domain.Loan) error {
	_, err := l.db.Exec(ctx, query.QueryUpdateLoan,
		loan.Code,
		loan.Principal,
		loan.InterestRate,
		loan.TotalWeeks,
		loan.WeeklyRepayment,
		loan.OutstandingBalance,
		loan.MissedPayments,
		loan.IsDelinquent,
		loan.IsActive,
		loan.OperatedBy,
		loan.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update loan: %v", err)
	}

	return nil
}

func (l *LoanRepository) CheckIsDelinquent(ctx context.Context, loanID int64) (bool, error) {
	loan, err := l.GetLoanByID(ctx, loanID)
	if err != nil {
		return false, err
	}

	if loan == nil {
		return false, helper.NewErrorx(http.StatusNotFound, "loan not found")
	}

	return loan.IsDelinquent, nil
}

func (l *LoanRepository) GetOutstandingBalance(ctx context.Context, loanID int64) (float64, error) {
	loan, err := l.GetLoanByID(ctx, loanID)
	if err != nil {
		return 0, err
	}

	if loan == nil {
		return 0, helper.NewErrorx(http.StatusNotFound, "loan not found")
	}

	return loan.OutstandingBalance, nil
}

func (l *LoanRepository) GetLoanByCode(ctx context.Context, code string) (*domain.Loan, error) {
	loan := &domain.Loan{}
	err := l.db.QueryRow(ctx, query.QueryGetLoanByCode, code).Scan(
		&loan.ID,
		&loan.Code,
		&loan.UserID,
		&loan.Principal,
		&loan.InterestRate,
		&loan.TotalWeeks,
		&loan.WeeklyRepayment,
		&loan.OutstandingBalance,
		&loan.MissedPayments,
		&loan.IsDelinquent,
		&loan.IsActive,
		&loan.CreatedBy,
		&loan.CreatedAt,
		&loan.UpdatedBy,
		&loan.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return loan, nil
}
