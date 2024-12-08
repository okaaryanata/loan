package service

import (
	"context"
	"log"
	"net/http"

	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/helper"
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

func (r *RepaymentService) CreateRepayment(ctx context.Context, repayment *domain.LoanRepayment) helper.Errorx {
	repayment.OperatedBy = helper.Chains(repayment.OperatedBy, "SYSTEM")
	err := r.repaymentRepo.CreateRepayment(ctx, nil, repayment)
	if err != nil {
		return helper.NewErrorxFromErr(err)
	}
	return nil
}

func (r *RepaymentService) GetRepaymentByIDAndLoanID(ctx context.Context, repaymentID, loanID int64) (*domain.LoanRepayment, helper.Errorx) {
	repayment, err := r.repaymentRepo.GetRepaymentByIDAndLoanID(ctx, repaymentID, loanID)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	return repayment, nil
}

func (r *RepaymentService) GetRepaymentByID(ctx context.Context, repaymentID int64) (*domain.LoanRepayment, helper.Errorx) {
	repayment, err := r.repaymentRepo.GetRepaymentByID(ctx, repaymentID)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}
	return repayment, nil
}

func (r *RepaymentService) GetRepaymentsByLoanID(ctx context.Context, loanID int64) ([]*domain.LoanRepayment, helper.Errorx) {
	repayments, err := r.repaymentRepo.GetRepaymentsByLoanID(ctx, loanID)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}
	return repayments, nil
}

func (r *RepaymentService) MakePayment(ctx context.Context, req *domain.MakePaymentRequest) helper.Errorx {
	req.OperatedBy = helper.Chains(req.OperatedBy, "SYSTEM")

	// Get Loan Detail
	loan, errx := r.loanService.GetLoanByIDandUserID(ctx, req.LoanID, req.UserID)
	if errx != nil {
		return errx
	}

	if loan == nil {
		return helper.NewErrorx(http.StatusNotFound, "loan not found")
	}

	if req.Amount != loan.WeeklyRepayment {
		return helper.NewErrorx(http.StatusBadRequest, "payment amount must be equal to weekly repayment")
	}

	// Get Repayment
	repayment, errx := r.GetRepaymentByIDAndLoanID(ctx, req.RepaymentID, loan.ID)
	if errx != nil {
		return errx
	}

	if repayment == nil {
		return helper.NewErrorx(http.StatusNotFound, "repayment not found")
	}

	if repayment.Paid {
		return helper.NewErrorx(http.StatusBadRequest, "payment for this week is already made")
	}

	// Make Repayment
	req.Paid = true
	err := r.repaymentRepo.MakePayment(ctx, req)
	if err != nil {
		return helper.NewErrorxFromErr(err)
	}

	// Update Loan & deliquent (async)
	go func() {
		var errUpdateLoan helper.Errorx
		childCtx := context.WithoutCancel(ctx)
		loan.UpdatedBy = req.OperatedBy
		loan.IsDelinquent, loan.MissedPayments, errUpdateLoan = r.loanService.CheckIsDelinquent(childCtx, loan.ID)
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

func (r *RepaymentService) PrintSchedule(ctx context.Context, userID int64, loanID int64) ([]*domain.LoanRepayment, helper.Errorx) {
	var (
		repayments []*domain.LoanRepayment
		err        error
	)
	if loanID != 0 {
		repayments, err = r.repaymentRepo.GetRepaymentsByLoanIDAndUserID(ctx, loanID, userID)
		if err != nil {
			return nil, helper.NewErrorxFromErr(err)
		}

		if len(repayments) == 0 {
			return nil, helper.NewErrorx(http.StatusNotFound, "repayments not found")
		}

		return repayments, nil
	}

	repayments, err = r.repaymentRepo.GetRepaymentsByUserID(ctx, userID)
	if err != nil {
		return nil, helper.NewErrorxFromErr(err)
	}

	return repayments, nil
}
