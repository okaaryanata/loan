package query

var (
	QueryCreateUser = `
		INSERT INTO loans.users (
			username, 
			is_active,
			created_by,
			updated_by
			) VALUES ($1, $2, $3, $3) RETURNING user_id
	`

	QueryGetUserByID = `
		SELECT
			user_id,
			username,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.users
		WHERE 
			user_id = $1
			AND is_active = true
	`

	QueryGetUserByUsernames = `
		SELECT
			user_id,
			username,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.users 
		WHERE 
			username = ANY($1)
			AND is_active = true
	`

	QueryGetUsers = `
		SELECT
			user_id,
			username,
			is_active,
			created_by,
			created_at,
			updated_by,
			updated_at
		FROM loans.users 
		WHERE 
			is_active = true
	`
)
