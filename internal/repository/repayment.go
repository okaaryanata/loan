package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okaaryanata/loan/internal/domain"
	"github.com/okaaryanata/loan/internal/repository/query"
)

type (
	RepaymentRepository struct {
		db *pgxpool.Pool
	}
)

func NewRepaymentRepository(db *pgxpool.Pool) *RepaymentRepository {
	return &RepaymentRepository{
		db: db,
	}
}

func (r *RepaymentRepository) CreateRepayment(ctx context.Context, tx pgx.Tx, repayment *domain.LoanRepayment) error {
	var err error
	if tx != nil {
		err = tx.QueryRow(ctx, query.QueryCreateRepayment,
			repayment.LoanID,
			repayment.Week,
			repayment.Amount,
			repayment.Paid,
			repayment.DueDate,
			repayment.IsActive,
			repayment.CreatedBy,
			repayment.UpdatedBy).Scan(&repayment.ID)
	} else {
		err = r.db.QueryRow(ctx, query.QueryCreateRepayment,
			repayment.LoanID,
			repayment.Week,
			repayment.Amount,
			repayment.Paid,
			repayment.DueDate,
			repayment.IsActive,
			repayment.CreatedBy,
			repayment.UpdatedBy).Scan(&repayment.ID)
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *RepaymentRepository) GetRepaymentByID(ctx context.Context, repaymentID int64) (*domain.LoanRepayment, error) {

	return nil, nil
}

func (r *RepaymentRepository) GetRepaymentsByLoanID(ctx context.Context, loanID int64) ([]*domain.LoanRepayment, error) {
	return nil, nil
}

func (r *RepaymentRepository) MakePayment(ctx context.Context, req *domain.MakePaymentRequest) error {
	return nil
}

func (r *RepaymentRepository) PrintSchedule(ctx context.Context, userID int64, loanID int64) error {
	return nil
}
