package repo

import (
	"context"
	"database/sql"
	"log"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
)

// GetUserDetail retrieves the user details from the database based on the provided email.
// It takes a *sql.DB object and the email string as parameters.
// It returns a models.User object and an error.
// The function queries the database for the user with the specified email and scans the result into a models.User object.
// If the query is successful, it returns the user object and nil error.
// If the query fails or no user is found, it returns an empty user object and the corresponding error.
func GetUserDetail(db *sql.DB, email string) (models.User, error) {
	row := db.QueryRow("SELECT id, username, email, phone, hidden_phone_number, fullname, hidden_email, avatar, gender, password_hash, two_factor_enabled, is_active, created_at, updated_at FROM users "+
		"WHERE email = $1 LIMIT 1", email)

	var i models.User

	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.HiddenPhoneNumber,
		&i.FullName,
		&i.HiddenEmail,
		&i.Avatar,
		&i.Gender,
		&i.PasswordHash,
		&i.TwoFactorEnabled,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

// CreateUser creates a new user in the database with the given email.
// It returns the created user and any error encountered.
func CreateUser(db *sql.DB, email string) (models.User, error) {
	row := db.QueryRow("INSERT INTO users (email) VALUES ($1) RETURNING id", email)
	var i models.User
	err := row.Scan(
		&i.ID,
	)
	return i, err
}

// UpdatePassword updates the password hash for a user in the database.
// It takes a database connection (`db`) and an `UpdatePasswordParams` struct (`arg`) as input.
// It returns an error if the update operation fails.
const updatePassword = `
UPDATE users
SET password_hash = $1, hidden_email = $3
WHERE id = $2
RETURNING id, email, hidden_email
`

func UpdatePassword(db *sql.DB, arg models.UpdatePasswordParams) (models.UpdateUserResponse, error) {
	var i models.UpdateUserResponse
	err := db.QueryRowContext(context.Background(), updatePassword, arg.PasswordHash, arg.ID, arg.HiddenEmail).Scan(&i.Id, &i.Email, &i.HiddenEmail)
	return i, err
}

const joinUsersWithVerificationByEmail = `-- name: JoinUsersWithVerificationByEmail :many
SELECT users.id, users.username, users.email, users.phone, users.hidden_phone_number, users.fullname, users.hidden_email, users.avatar, users.gender, users.password_hash, users.two_factor_enabled, users.is_active, users.created_at, users.updated_at
FROM users
JOIN verification ON users.id = verification.user_id
WHERE verification.is_verified = true
AND users.email = $1
`

func JoinUsersWithVerificationByEmail(db *sql.DB, email string) ([]models.User, error) {
	rows, err := db.QueryContext(context.Background(), joinUsersWithVerificationByEmail, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.User{}
	for rows.Next() {
		var i models.User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Phone,
			&i.HiddenPhoneNumber,
			&i.FullName,
			&i.HiddenEmail,
			&i.Avatar,
			&i.Gender,
			&i.PasswordHash,
			&i.TwoFactorEnabled,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	log.Print(items)
	return items, nil
}

const joinUsersWithVerificationByPhone = `-- name: JoinUsersWithVerificationByPhone :many
SELECT users.id, users.username, users.email, users.phone, users.hidden_phone_number, users.fullname, users.hidden_email, users.avatar, users.gender, users.password_hash, users.two_factor_enabled, users.is_active, users.created_at, users.updated_at
FROM users
JOIN verification ON users.id = verification.user_id
WHERE verification.is_verified = true
AND users.phone = $1
`

func JoinUsersWithVerificationByPhone(db *sql.DB, phone string) ([]models.User, error) {
	rows, err := db.QueryContext(context.Background(), joinUsersWithVerificationByPhone, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.User{}
	for rows.Next() {
		var i models.User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Phone,
			&i.HiddenPhoneNumber,
			&i.FullName,
			&i.HiddenEmail,
			&i.Avatar,
			&i.Gender,
			&i.PasswordHash,
			&i.TwoFactorEnabled,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const joinUsersWithVerificationByUsername = `-- name: JoinUsersWithVerificationByUsername :many
SELECT users.id, users.username, users.email, users.phone, users.hidden_phone_number, users.fullname, users.hidden_email, users.avatar, users.gender, users.password_hash, users.two_factor_enabled, users.is_active, users.created_at, users.updated_at
FROM users
JOIN verification ON users.id = verification.user_id
WHERE verification.is_verified = true
AND users.username = $1
`

func JoinUsersWithVerificationByUsername(db *sql.DB, username string) ([]models.User, error) {
	rows, err := db.QueryContext(context.Background(), joinUsersWithVerificationByUsername, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.User{}
	for rows.Next() {
		var i models.User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Phone,
			&i.HiddenPhoneNumber,
			&i.FullName,
			&i.HiddenEmail,
			&i.Avatar,
			&i.Gender,
			&i.PasswordHash,
			&i.TwoFactorEnabled,
			&i.IsActive,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}