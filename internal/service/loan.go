package service

import (
	"context"
	"net/http"

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

func (l *LoanService) CreateLoan(ctx context.Context, req *domain.LoanRequest) (*domain.Loan, helper.Errorx) {
	// Check Loan Code
	loan, err := l.loanRepo.GetLoanByCode(ctx, req.Code)
	if err == nil && loan != nil {
		return nil, helper.NewErrorxif(http.StatusBadRequest, "loan with code %s already exists", req.Code)
	}

	// Main Calculation
	totalAmount := req.Principal + (req.Principal * req.InterestRate)
	weeklyRepayment := totalAmount / float64(req.TotalWeeks)

	req.OperatedBy = helper.Chains(req.OperatedBy, "SYSTEM")
	loan = &domain.Loan{
		UserID:             req.UserID,
		Code:               req.Code,
		Principal:          req.Principal,
		InterestRate:       req.InterestRate,
		TotalWeeks:         req.TotalWeeks,
		WeeklyRepayment:    weeklyRepayment,
		OutstandingBalance: totalAmount,
		IsActive:           true,
		OperatedBy:         req.OperatedBy,
	}

	tx, err := l.db.Begin(ctx)
	if err != nil {
		return nil, helper.NewErrorxf("failed to begin transaction: %v", err)
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
		return nil, helper.NewErrorxFromErr(err)
	}

	// Create Repayments schedule
	localTime, err := helper.JakartaTime()
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
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
			return nil, helper.NewErrorxFromErr(err)
		}
	}

	return loan, nil
}

func (l *LoanService) GetLoanByID(ctx context.Context, loanID int64) (*domain.Loan, helper.Errorx) {
	loan, err := l.loanRepo.GetLoanByID(ctx, loanID)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	return loan, nil
}

func (l *LoanService) GetLoanByCode(ctx context.Context, code string) (*domain.Loan, helper.Errorx) {
	loan, err := l.loanRepo.GetLoanByCode(ctx, code)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	if loan == nil {
		return nil, helper.NewErrorx(http.StatusNotFound, "loan not found")
	}

	return loan, nil
}

func (l *LoanService) GetLoansByUserID(ctx context.Context, userID int64) ([]*domain.Loan, helper.Errorx) {
	loans, err := l.loanRepo.GetLoansByUserID(ctx, userID)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	return loans, nil
}

func (l *LoanService) GetLoanByIDandUserID(ctx context.Context, loanID, userID int64) (*domain.Loan, helper.Errorx) {
	loan, err := l.loanRepo.GetLoansByIDandUserID(ctx, loanID, userID)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	return loan, nil
}

func (l *LoanService) CheckIsDelinquent(ctx context.Context, loanID int64, currentWeek int) (bool, int, helper.Errorx) {
	var (
		totalMissedPayments int
		isDelinquent        bool
	)

	repayments, err := l.repaymentRepo.GetRepaymentsByLoanID(ctx, loanID)
	if err != nil {
		return false, 0, helper.NewErrorxFromErr(err)
	}

	for idx := range repayments {
		if repayments[idx].Week > currentWeek {
			break
		}

		if !repayments[idx].Paid {
			totalMissedPayments++
		}
	}

	if totalMissedPayments >= 2 {
		isDelinquent = true
	}

	return isDelinquent, totalMissedPayments, nil
}

func (l *LoanService) GetOutstandingBalance(ctx context.Context, loanID int64) (float64, helper.Errorx) {
	outstanding, err := l.loanRepo.GetOutstandingBalance(ctx, loanID)
	if err != nil {
		return 0, helper.NewErrorxFromErr(err)
	}

	return outstanding, nil
}

func (l *LoanService) UpdateLoan(ctx context.Context, loan *domain.Loan) helper.Errorx {
	loan.OperatedBy = helper.Chains(loan.OperatedBy, "SYSTEM")
	err := l.loanRepo.UpdateLoan(ctx, loan)
	if err != nil {
		return helper.NewErrorxFromErr(err)
	}

	return nil
}
