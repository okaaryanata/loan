package service

import (
	"context"
	"errors"
	"log"

	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/repository"
)

type (
	RepaymentService struct {
		// service
		loanService *LoanService

		// repository
		repaymentRepo *repository.RepaymentRepository
	}
)

func NewRepaymentService(
	loanService *LoanService,
	repaymentRepo *repository.RepaymentRepository,
) *RepaymentService {
	return &RepaymentService{
		// service
		loanService: loanService,
		// repository
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

func (r *RepaymentService) MakePayment(ctx context.Context, req *domain.MakePaymentRequest) error {
	// Get Loan Detail
	loan, err := r.loanService.GetLoanByID(ctx, req.LoanID)
	if err != nil {
		return err
	}

	if req.Week < 1 || req.Week > loan.TotalWeeks {
		return errors.New("invalid week number")
	}

	if req.Amount != loan.WeeklyRepayment {
		return errors.New("payment amount must be equal to weekly repayment")
	}

	// Get Repayment
	repayment, err := r.GetRepaymentByID(ctx, req.LoanID)
	if err != nil {
		return err
	}

	if repayment.Paid {
		return errors.New("payment for this week is already made")
	}

	// Make Repayment
	err = r.repaymentRepo.MakePayment(ctx, req)
	if err != nil {
		return err
	}

	// Update Loan & deliquent (async)
	go func() {
		var errUpdateLoan error
		childCtx := context.WithoutCancel(ctx)
		loan.IsDelinquent, loan.MissedPayments, errUpdateLoan = r.loanService.CheckIsDelinquent(childCtx, loan.LoanID, req.Week)
		if errUpdateLoan != nil {
			log.Println(errUpdateLoan)
			return
		}

		loan.OutstandingBalance -= req.Amount
		errUpdateLoan = r.loanService.UpdateLoan(childCtx, loan)
		if errUpdateLoan != nil {
			log.Println(errUpdateLoan)
			return
		}
	}()

	return nil
}

func (r *RepaymentService) PrintSchedule(ctx context.Context, userID int64, loanID int64) error {
	return r.repaymentRepo.PrintSchedule(ctx, userID, loanID)
}
