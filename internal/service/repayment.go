package service

import (
	"context"

	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/repository"
)

type (
	RepaymentService struct {
		repaymentRepo *repository.RepaymentRepository
	}
)

func NewRepaymentService(
	repaymentRepo *repository.RepaymentRepository,
) *RepaymentService {
	return &RepaymentService{
		repaymentRepo: repaymentRepo,
	}
}

func (r *RepaymentService) CreateRepayment(ctx context.Context, repayment *domain.LoanRepayment) (*domain.LoanRepayment, error) {
	return r.repaymentRepo.CreateRepayment(ctx, repayment)
}

func (r *RepaymentService) GetRepaymentByID(ctx context.Context, repaymentID int64) (*domain.LoanRepayment, error) {
	return r.repaymentRepo.GetRepaymentByID(ctx, repaymentID)
}

func (r *RepaymentService) GetRepaymentsByLoanID(ctx context.Context, loanID int64) ([]*domain.LoanRepayment, error) {
	return r.repaymentRepo.GetRepaymentsByLoanID(ctx, loanID)
}

func (r *RepaymentService) MakePayment(ctx context.Context, loanID int64, week int, amount int) error {
	return r.repaymentRepo.MakePayment(ctx, loanID, week, amount)
}

func (r *RepaymentService) PrintSchedule(ctx context.Context, loanID int64) error {
	return r.repaymentRepo.PrintSchedule(ctx, loanID)
}
