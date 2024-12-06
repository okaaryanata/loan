package service

import (
	"context"

	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/repository"
)

type (
	LoanService struct {
		loanRepo      *repository.LoanRepository
		repaymentRepo *repository.RepaymentRepository
	}
)

func NewLoanService(
	loanRepo *repository.LoanRepository,
	repaymentRepo *repository.RepaymentRepository,
) *LoanService {
	return &LoanService{
		loanRepo:      loanRepo,
		repaymentRepo: repaymentRepo,
	}
}

func (l *LoanService) CreateLoan(ctx context.Context, req *domain.LoanRequest) (*domain.Loan, error) {
	// Main Calculation
	totalAmount := req.Principal + (req.Principal * req.Interest)
	weeklyRepayment := totalAmount / float64(req.Weeks)

	loan := &domain.Loan{
		UserID:             req.UserID,
		Principal:          req.Principal,
		InterestRate:       req.Interest,
		TotalWeeks:         req.Weeks,
		WeeklyRepayment:    weeklyRepayment,
		OutstandingBalance: totalAmount,
	}

	// Create Loan
	err := l.loanRepo.CreateLoan(ctx, loan)
	if err != nil {
		return nil, err
	}

	// Create Repayments schedule
	for week := 1; week < req.Weeks+1; week++ {
		_, err := l.repaymentRepo.CreateRepayment(ctx, &domain.LoanRepayment{
			LoanID: loan.LoanID,
			Week:   week,
			Amount: weeklyRepayment,
		})
		if err != nil {
			return nil, err
		}
	}

	return loan, nil
}

func (l *LoanService) GetLoanByID(ctx context.Context, loanID int64) (*domain.Loan, error) {
	return l.loanRepo.GetLoanByID(ctx, loanID)
}

func (l *LoanService) GetLoansByUserID(ctx context.Context, userID int64) ([]*domain.Loan, error) {
	return l.loanRepo.GetLoansByUserID(ctx, userID)
}

func (l *LoanService) CheckIsDelinquent(ctx context.Context, loanID int64, currentWeek int) (bool, int, error) {
	var (
		totalMissedPayments int
		isDelinquent        bool
	)

	repayments, err := l.repaymentRepo.GetRepaymentsByLoanID(ctx, loanID)
	if err != nil {
		return false, 0, err
	}

	for idx := range repayments {
		if repayments[idx].Week > currentWeek {
			break
		}

		if !repayments[idx].Paid {
			totalMissedPayments++
			if totalMissedPayments >= 2 {
				isDelinquent = true
				break
			}
		}
	}

	if totalMissedPayments >= 2 {
		isDelinquent = true
	}

	return isDelinquent, totalMissedPayments, nil
}

func (l *LoanService) GetOutstandingBalance(ctx context.Context, loanID int64) (int, error) {
	return l.loanRepo.GetOutstandingBalance(ctx, loanID)
}

func (l *LoanService) UpdateLoan(ctx context.Context, loan *domain.Loan) error {
	return l.loanRepo.UpdateLoan(ctx, loan)
}
