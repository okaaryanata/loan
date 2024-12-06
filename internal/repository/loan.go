package repository

import (
	"context"
	"fmt"

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
	row := l.db.QueryRow(ctx, query.QueryGetLoanByID, loanID)
	loan := &domain.Loan{}
	err := helper.StructScan(row, loan)
	if err != nil {
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
	err = helper.StructScanAll(rows, loans)
	if err != nil {
		return nil, err
	}

	return loans, nil
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

	return loan.IsDelinquent, nil
}

func (l *LoanRepository) GetOutstandingBalance(ctx context.Context, loanID int64) (float64, error) {
	loan, err := l.GetLoanByID(ctx, loanID)
	if err != nil {
		return 0, err
	}

	return loan.OutstandingBalance, nil
}
