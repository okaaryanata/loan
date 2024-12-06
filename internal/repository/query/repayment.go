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
		SELECT * FROM loans.loan_repayments 
		WHERE 
			loan_repayment_id = $1
			and is_active = true;
	`

	QueryGetRepaymentsByLoanID = `
		SELECT * FROM loans.loan_repayments 
		WHERE 
			loan_id = $1
			and is_active = true;
	`
)
