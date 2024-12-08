package query

var (
	QueryCreateRepayment = `
		INSERT INTO loans.loan_repayments (
			loan_id, 
			week, 
			amount, 
			paid, 
			due_date, 
			is_active, 
			created_by, 
			updated_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING loan_repayment_id
	`

	QueryMakePayment = `
		UPDATE loans.loan_repayments SET
			paid = $1,
			updated_by = $2,
			updated_at = NOW()
		WHERE 
			loan_repayment_id = $3
			and is_active = true;
	`

	QueryGetRepaymentByID = `
		SELECT
			loan_repayment_id,
			loan_id,
			week,
			amount,
			paid,
			due_date,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.loan_repayments 
		WHERE 
			loan_repayment_id = $1
			and is_active = true;
	`

	QueryGetRepaymentByIDAndLoanID = `
		SELECT
			loan_repayment_id,
			loan_id,
			week,
			amount,
			paid,
			due_date,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.loan_repayments 
		WHERE 
			loan_repayment_id = $1
			and loan_id = $2
			and is_active = true;
	`

	QueryGetRepaymentsByLoanID = `
		SELECT 
			loan_repayment_id,
			loan_id,
			week,
			amount,
			paid,
			due_date,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.loan_repayments 
		WHERE 
			loan_id = $1
			and is_active = true
		ORDER BY due_date ASC;
	`

	QueryGetRepaymentsByLoanIDAndUserID = `
		SELECT 
			r.loan_repayment_id,
			r.loan_id,
			r.week,
			r.amount,
			r.paid,
			r.due_date,
			r.is_active,
			r.created_by,
			r.created_at,
			r.updated_by,
			r.updated_at
		FROM loans.loan_repayments r
		JOIN loans.loans l ON r.loan_id = l.loan_id
		WHERE 
			r.loan_id = $1
			and l.user_id = $2
			and r.is_active = true
		ORDER BY r.due_date ASC;
	`

	QueryGetRepaymentsByUserID = `
		SELECT 
			r.loan_repayment_id,
			r.loan_id,
			r.week,
			r.amount,
			r.paid,
			r.due_date,
			r.is_active,
			r.created_by,
			r.created_at,
			r.updated_by,
			r.updated_at
		FROM loans.loan_repayments r
		JOIN loans.loans l ON r.loan_id = l.loan_id
		WHERE 
			l.user_id = $1
			and r.is_active = true
		ORDER BY r.due_date ASC;
	`

	QueryGetLastPaidRepaymentByLoanID = `
		SELECT 
			r.loan_repayment_id,
			r.loan_id,
			r.week,
			r.amount,
			r.paid,
			r.due_date,
			r.is_active,
			r.created_by,
			r.created_at,
			r.updated_by,
			r.updated_at
		FROM loans.loan_repayments r
		WHERE 
			r.loan_id = $1
			and r.paid = true
			and r.is_active = true
		ORDER BY r.due_date DESC
		LIMIT 1;
	`
)
