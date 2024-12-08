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
	repayment := &domain.LoanRepayment{}
	err := r.db.QueryRow(ctx, query.QueryGetRepaymentByID, repaymentID).Scan(
		&repayment.ID,
		&repayment.LoanID,
		&repayment.Week,
		&repayment.Amount,
		&repayment.Paid,
		&repayment.DueDate,
		&repayment.IsActive,
		&repayment.CreatedBy,
		&repayment.CreatedAt,
		&repayment.UpdatedBy,
		&repayment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return repayment, nil
}

func (r *RepaymentRepository) GetRepaymentByIDAndLoanID(ctx context.Context, repaymentID, loanID int64) (*domain.LoanRepayment, error) {
	repayment := &domain.LoanRepayment{}
	err := r.db.QueryRow(ctx, query.QueryGetRepaymentByIDAndLoanID, repaymentID, loanID).Scan(
		&repayment.ID,
		&repayment.LoanID,
		&repayment.Week,
		&repayment.Amount,
		&repayment.Paid,
		&repayment.DueDate,
		&repayment.IsActive,
		&repayment.CreatedBy,
		&repayment.CreatedAt,
		&repayment.UpdatedBy,
		&repayment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return repayment, nil
}

func (r *RepaymentRepository) GetRepaymentsByLoanID(ctx context.Context, loanID int64) ([]*domain.LoanRepayment, error) {
	var repayments []*domain.LoanRepayment
	rows, err := r.db.Query(ctx, query.QueryGetRepaymentsByLoanID, loanID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var repayment domain.LoanRepayment
		err := rows.Scan(
			&repayment.ID,
			&repayment.LoanID,
			&repayment.Week,
			&repayment.Amount,
			&repayment.Paid,
			&repayment.DueDate,
			&repayment.IsActive,
			&repayment.CreatedBy,
			&repayment.CreatedAt,
			&repayment.UpdatedBy,
			&repayment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		repayments = append(repayments, &repayment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return repayments, nil
}

func (r *RepaymentRepository) GetRepaymentsByLoanIDAndUserID(ctx context.Context, loanID, userID int64) ([]*domain.LoanRepayment, error) {
	var repayments []*domain.LoanRepayment
	rows, err := r.db.Query(ctx, query.QueryGetRepaymentsByLoanIDAndUserID, loanID, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var repayment domain.LoanRepayment
		err := rows.Scan(
			&repayment.ID,
			&repayment.LoanID,
			&repayment.Week,
			&repayment.Amount,
			&repayment.Paid,
			&repayment.DueDate,
			&repayment.IsActive,
			&repayment.CreatedBy,
			&repayment.CreatedAt,
			&repayment.UpdatedBy,
			&repayment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		repayments = append(repayments, &repayment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return repayments, nil
}

func (r *RepaymentRepository) GetRepaymentsByUserID(ctx context.Context, userID int64) ([]*domain.LoanRepayment, error) {
	var repayments []*domain.LoanRepayment
	rows, err := r.db.Query(ctx, query.QueryGetRepaymentsByUserID, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var repayment domain.LoanRepayment
		err := rows.Scan(
			&repayment.ID,
			&repayment.LoanID,
			&repayment.Week,
			&repayment.Amount,
			&repayment.Paid,
			&repayment.DueDate,
			&repayment.IsActive,
			&repayment.CreatedBy,
			&repayment.CreatedAt,
			&repayment.UpdatedBy,
			&repayment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		repayments = append(repayments, &repayment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return repayments, nil
}

func (r *RepaymentRepository) MakePayment(ctx context.Context, req *domain.MakePaymentRequest) error {
	_, err := r.db.Exec(ctx, query.QueryMakePayment,
		req.Paid,
		req.OperatedBy,
		req.RepaymentID,
	)
	if err != nil {
		return err
	}

	return nil
}
