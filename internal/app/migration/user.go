package migration

const (
	QueryInitTableUsers = `
	CREATE TABLE IF NOT EXISTS loans.users (
	    user_id BIGSERIAL PRIMARY KEY,
	    username VARCHAR(50) NOT NULL UNIQUE,
	    is_active BOOLEAN DEFAULT TRUE,
	    created_at TIMESTAMP DEFAULT NOW(),
	    created_by VARCHAR(255) DEFAULT 'SYSTEM',
	    updated_at TIMESTAMP DEFAULT NOW(),
	    updated_by VARCHAR(255) DEFAULT 'SYSTEM',
	    deleted_at TIMESTAMP,
	    deleted_by VARCHAR(255)
	);

	-- Add indexes
	CREATE INDEX IF NOT EXISTS idx_users_username ON loans.users (username);
	CREATE INDEX IF NOT EXISTS idx_users_is_active ON loans.users (is_active);
	CREATE INDEX IF NOT EXISTS idx_users_created_at ON loans.users (created_at);
	CREATE INDEX IF NOT EXISTS idx_users_updated_at ON loans.users (updated_at);
	CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON loans.users (deleted_at);
	`
)
