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
		SELECT * FROM loans.users
		WHERE 
			user_id = $1
			AND is_active = true
	`

	QueryGetUserByUsername = `
		SELECT * FROM loans.users 
		WHERE 
			username = $1
			AND is_active = true
	`
)
