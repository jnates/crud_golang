package db

const (
	QueryGetUserByID = `
		SELECT id, name, email
		FROM users
		WHERE id = $1
	`

	QueryInsertUser = `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id
	`

	QueryUpdateUser = `
		UPDATE users
		SET name = $1, email = $2
		WHERE id = $3
	`

	QueryDeleteUser = `
		DELETE FROM users
		WHERE id = $1
	`

	QuerySelectUserBase = `
		SELECT id, name, email
		FROM users
	`
)
