package migration

const (
	QueryInitTableLoans = `
		CREATE TABLE IF NOT EXISTS loans.loans (
	    loan_id BIGSERIAL PRIMARY KEY,
	    user_id BIGINT NOT NULL,
	    code VARCHAR(8) NOT NULL,
	    principal NUMERIC(15, 2) NOT NULL,
	    interest_rate NUMERIC(3, 2) NOT NULL,
	    total_weeks INT NOT NULL,
	    weekly_repayment NUMERIC(15, 2) NOT NULL,
	    outstanding_balance NUMERIC(15, 2) DEFAULT 0.0,
	    missed_payments INT DEFAULT 0,
	    is_delinquent BOOLEAN DEFAULT FALSE,
	    is_active BOOLEAN DEFAULT TRUE,
	    created_at TIMESTAMP DEFAULT NOW(),
	    created_by VARCHAR(255) DEFAULT 'SYSTEM',
	    updated_at TIMESTAMP DEFAULT NOW(),
	    updated_by VARCHAR(255) DEFAULT 'SYSTEM',
	    deleted_at TIMESTAMP,
	    deleted_by VARCHAR(255),
	    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES loans.users (user_id) ON DELETE CASCADE
	);

	-- Add indexes
	CREATE INDEX IF NOT EXISTS idx_loans_code ON loans.loans (code);
	CREATE INDEX IF NOT EXISTS idx_loans_principal ON loans.loans (principal);
	CREATE INDEX IF NOT EXISTS idx_loans_interest_rate ON loans.loans (interest_rate);
	CREATE INDEX IF NOT EXISTS idx_loans_total_weeks ON loans.loans (total_weeks);
	CREATE INDEX IF NOT EXISTS idx_loans_outstanding_balance ON loans.loans (outstanding_balance);
	CREATE INDEX IF NOT EXISTS idx_loans_missed_payments ON loans.loans (missed_payments);
	CREATE INDEX IF NOT EXISTS idx_loans_is_delinquent ON loans.loans (is_delinquent);
	CREATE INDEX IF NOT EXISTS idx_loans_is_active ON loans.loans (is_active);
	CREATE INDEX IF NOT EXISTS idx_loans_created_at ON loans.loans (created_at);
	CREATE INDEX IF NOT EXISTS idx_loans_updated_at ON loans.loans (updated_at);
	`
)
