package repository

import (
	"context"

	"github.com/okaaryanata/loan/internal/domain"
)

type (
	RepaymentRepository struct{}
)

func NewRepaymentRepository() *RepaymentRepository {
	return &RepaymentRepository{}
}

func (r *RepaymentRepository) CreateRepayment(ctx context.Context, repayment *domain.LoanRepayment) (*domain.LoanRepayment, error) {
	return nil, nil
}

func (r *RepaymentRepository) GetRepaymentByID(ctx context.Context, repaymentID int64) (*domain.LoanRepayment, error) {
	return nil, nil
}

func (r *RepaymentRepository) GetRepaymentsByLoanID(ctx context.Context, loanID int64) ([]*domain.LoanRepayment, error) {
	return nil, nil
}

func (r *RepaymentRepository) MakePayment(ctx context.Context, loanID int64, week int, amount int) error {
	return nil
}

func (r *RepaymentRepository) PrintSchedule(ctx context.Context, loanID int64) error {
	return nil
}
