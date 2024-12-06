package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/helper"
	"github.com/okaaryanata/loan/internal/repository"
)

type (
	LoanService struct {
		db            *pgxpool.Pool
		loanRepo      *repository.LoanRepository
		repaymentRepo *repository.RepaymentRepository
	}
)

func NewLoanService(
	db *pgxpool.Pool,
	loanRepo *repository.LoanRepository,
	repaymentRepo *repository.RepaymentRepository,
) *LoanService {
	return &LoanService{
		db:            db,
		loanRepo:      loanRepo,
		repaymentRepo: repaymentRepo,
	}
}

func (l *LoanService) CreateLoan(ctx context.Context, req *domain.LoanRequest) (*domain.Loan, error) {
	// Main Calculation
	totalAmount := req.Principal + (req.Principal * req.InterestRate)
	weeklyRepayment := totalAmount / float64(req.TotalWeeks)

	loan := &domain.Loan{
		UserID:             req.UserID,
		Principal:          req.Principal,
		InterestRate:       req.InterestRate,
		TotalWeeks:         req.TotalWeeks,
		WeeklyRepayment:    weeklyRepayment,
		OutstandingBalance: totalAmount,
	}

	tx, err := l.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	// Create Loan
	err = l.loanRepo.CreateLoan(ctx, tx, loan)
	if err != nil {
		return nil, err
	}

	// Create Repayments schedule
	localTime, err := helper.JakartaTime()
	if err != nil {
		return nil, err
	}
	for week := 1; week < req.TotalWeeks+1; week++ {
		*localTime = localTime.AddDate(0, 0, 7)
		err := l.repaymentRepo.CreateRepayment(ctx, tx, &domain.LoanRepayment{
			LoanID:     loan.ID,
			Week:       week,
			Amount:     weeklyRepayment,
			DueDate:    *localTime,
			IsActive:   true,
			OperatedBy: req.OperatedBy,
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

func (l *LoanService) GetOutstandingBalance(ctx context.Context, loanID int64) (float64, error) {
	return l.loanRepo.GetOutstandingBalance(ctx, loanID)
}

func (l *LoanService) UpdateLoan(ctx context.Context, loan *domain.Loan) error {
	return l.loanRepo.UpdateLoan(ctx, loan)
}
