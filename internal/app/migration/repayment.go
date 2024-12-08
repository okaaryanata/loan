package migration

const (
	QueryInitTableRepayments = `
	CREATE TABLE IF NOT EXISTS loans.loan_repayments (
	    loan_repayment_id BIGSERIAL PRIMARY KEY,
	    loan_id BIGINT NOT NULL,
	    week INT NOT NULL,
	    amount NUMERIC(15, 2) NOT NULL,
	    paid BOOLEAN DEFAULT FALSE,
	    due_date TIMESTAMP NOT NULL,
	    is_active BOOLEAN DEFAULT TRUE,
	    created_at TIMESTAMP DEFAULT NOW(),
	    created_by VARCHAR(255) DEFAULT 'SYSTEM',
	    updated_at TIMESTAMP DEFAULT NOW(),
	    updated_by VARCHAR(255) DEFAULT 'SYSTEM',
	    deleted_at TIMESTAMP,
	    deleted_by VARCHAR(255),
	    CONSTRAINT fk_loan FOREIGN KEY (loan_id) REFERENCES loans.loans (loan_id) ON DELETE CASCADE
	);

	-- Add indexes
	CREATE INDEX IF NOT EXISTS idx_loan_repayments_loan_id ON loans.loan_repayments (loan_id);
	CREATE INDEX IF NOT EXISTS idx_loan_repayments_week ON loans.loan_repayments (week);
	CREATE INDEX IF NOT EXISTS idx_loan_repayments_amount ON loans.loan_repayments (amount);
	CREATE INDEX IF NOT EXISTS idx_loan_repayments_paid ON loans.loan_repayments (paid);
	CREATE INDEX IF NOT EXISTS idx_loan_repayments_due_date ON loans.loan_repayments (due_date);
	CREATE INDEX IF NOT EXISTS idx_loan_repayments_is_active ON loans.loan_repayments (is_active);
	CREATE INDEX IF NOT EXISTS idx_loan_repayments_created_at ON loans.loan_repayments (created_at);
	CREATE INDEX IF NOT EXISTS idx_loan_repayments_updated_at ON loans.loan_repayments (updated_at);
	CREATE INDEX IF NOT EXISTS idx_loan_repayments_deleted_at ON loans.loan_repayments (deleted_at);
	`
)
