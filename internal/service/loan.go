package service

import (
	"context"

	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/repository"
)

type (
	LoanService struct {
		loanRepo *repository.LoanRepository
	}
)

func NewLoanService(
	loanRepo *repository.LoanRepository,
) *LoanService {
	return &LoanService{
		loanRepo: loanRepo,
	}
}

func (l *LoanService) CreateLoan(ctx context.Context, loan *domain.Loan) (*domain.Loan, error) {
	return l.loanRepo.CreateLoan(ctx, loan)
}

func (l *LoanService) GetLoanByID(ctx context.Context, loanID int64) (*domain.Loan, error) {
	return l.loanRepo.GetLoanByID(ctx, loanID)
}

func (l *LoanService) GetLoansByUserID(ctx context.Context, userID int64) ([]*domain.Loan, error) {
	return l.loanRepo.GetLoansByUserID(ctx, userID)
}

func (l *LoanService) CheckIsDelinquent(ctx context.Context, loanID int64) (bool, error) {
	return l.loanRepo.CheckIsDelinquent(ctx, loanID)
}

func (l *LoanService) GetOutstandingBalance(ctx context.Context, loanID int64) (int, error) {
	return l.loanRepo.GetOutstandingBalance(ctx, loanID)
}
