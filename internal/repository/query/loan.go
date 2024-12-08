package query

var (
	QueryCreateLoan = `
		INSERT INTO loans.loans (
			user_id,
			code, 
			principal, 
			interest_rate, 
			total_weeks, 
			weekly_repayment, 
			outstanding_balance, 
			missed_payments, 
			is_delinquent, 
			is_active, 
			created_by,
			updated_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $11) 
		RETURNING loan_id;
	`

	QueryUpdateLoan = `
		UPDATE loans.loans SET
			code = $1, 
			principal = $2, 
			interest_rate = $3, 
			total_weeks = $4, 
			weekly_repayment = $5, 
			outstanding_balance = $6, 
			missed_payments = $7, 
			is_delinquent = $8, 
			is_active = $9, 
			updated_by = $10,
			updated_at = NOW()
		WHERE 
			loan_id = $11
			and is_active = true;
	`

	QueryGetLoanByID = `
		SELECT
			loan_id,
			code,
			user_id,
			principal,
			interest_rate,
			total_weeks,
			weekly_repayment,
			outstanding_balance,
			missed_payments,
			is_delinquent,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.loans 
		WHERE 
			loan_id = $1
			and is_active = true;
	`

	QueryGetLoanByCode = `
		SELECT
			loan_id,
			code,
			user_id,
			principal,
			interest_rate,
			total_weeks,
			weekly_repayment,
			outstanding_balance,
			missed_payments,
			is_delinquent,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.loans 
		WHERE 
			code = $1
			and is_active = true;
	`

	QueryGetLoansByUserID = `
		SELECT 
			loan_id,
			code,
			user_id,
			principal,
			interest_rate,
			total_weeks,
			weekly_repayment,
			outstanding_balance,
			missed_payments,
			is_delinquent,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.loans 
		WHERE 
			user_id = $1
			and is_active = true;
	`

	QueryGetLoanByIDandUserID = `
		SELECT
			loan_id,
			code,
			user_id,
			principal,
			interest_rate,
			total_weeks,
			weekly_repayment,
			outstanding_balance,
			missed_payments,
			is_delinquent,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.loans 
		WHERE 
			loan_id = $1
			and user_id = $2
			and is_active = true;
	`
)
