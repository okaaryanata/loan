package repository

import (
	"context"

	"github.com/okaaryanata/loan/internal/domain"
)

type (
	LoanRepository struct{}
)

func NewLoanRepository() *LoanRepository {
	return &LoanRepository{}
}

func (l *LoanRepository) CreateLoan(ctx context.Context, loan *domain.Loan) (*domain.Loan, error) {
	return nil, nil
}

func (l *LoanRepository) GetLoanByID(ctx context.Context, loanID int64) (*domain.Loan, error) {
	return nil, nil
}

func (l *LoanRepository) GetLoansByUserID(ctx context.Context, userID int64) ([]*domain.Loan, error) {
	return nil, nil
}

func (l *LoanRepository) CheckIsDelinquent(ctx context.Context, loanID int64) (bool, error) {
	return false, nil
}

func (l *LoanRepository) GetOutstandingBalance(ctx context.Context, loanID int64) (int, error) {
	return 0, nil
}
