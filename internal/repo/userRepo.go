package repo

import (
	"context"
	"database/sql"
	"fmt"

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
SET password_hash = $1, hidden_email = $2, is_active = true
WHERE id = $3
RETURNING id, email, hidden_email, is_active
`

func UpdatePassword(db *sql.DB, arg models.UpdatePasswordParams) (models.UpdateUserResponse, error) {
	var i models.UpdateUserResponse
	err := db.QueryRowContext(context.Background(), updatePassword, arg.PasswordHash, arg.HiddenEmail, arg.ID).Scan(&i.Id, &i.Email, &i.HiddenEmail, &i.IsActive)
	return i, err
}

const joinUsersWithVerificationByEmail = `-- name: JoinUsersWithVerificationByEmail :many
SELECT users.id, users.username, users.email, users.phone, users.hidden_phone_number, users.fullname, users.hidden_email, users.avatar, users.gender, users.password_hash, users.two_factor_enabled, users.is_active, users.created_at, users.updated_at
FROM users
JOIN verification ON users.id = verification.user_id
WHERE verification.is_verified = true
AND users.email = $1
`

// JoinUsersWithVerificationByEmail joins the user table with the verification table based on the provided email.
// It returns a slice of User models and an error if any occurred.
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
	return items, nil
}

const joinUsersWithVerificationByPhone = `-- name: JoinUsersWithVerificationByPhone :many
SELECT users.id, users.username, users.email, users.phone, users.hidden_phone_number, users.fullname, users.hidden_email, users.avatar, users.gender, users.password_hash, users.two_factor_enabled, users.is_active, users.created_at, users.updated_at
FROM users
JOIN verification ON users.id = verification.user_id
WHERE verification.is_verified = true
AND users.phone = $1
`

// JoinUsersWithVerificationByPhone retrieves a list of users along with their verification information
// based on the provided phone number.
// It takes a database connection `db` and a `phone` string as input parameters.
// It returns a slice of `models.User` and an error if any.
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

// JoinUsersWithVerificationByUsername joins the user table with the verification table based on the provided username.
// It returns a slice of models.User and an error if any.
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

const updateOnlyPassword = `-- name: UpdatePassword :exec
UPDATE users
SET password_hash = $1
WHERE id = $2
`

// UpdateOnlyPassword updates the password hash for a user in the database.
// It takes a database connection (`db`) and an argument (`arg`) of type `models.UpdateOnlyPasswordParams`.
// It returns an error if the update operation fails.
func UpdateOnlyPassword(db *sql.DB, arg models.UpdateOnlyPasswordParams) error {
	_, err := db.ExecContext(context.Background(), updateOnlyPassword, arg.PasswordHash, arg.ID)
	return err
}

const getUserId = `-- name: GetUserId :one
SELECT id, username, email, phone, hidden_phone_number, fullname, hidden_email, avatar, gender, two_factor_enabled, is_active, created_at FROM users
WHERE id = $1 AND is_active = $2 LIMIT 1
`

func GetUserId(db *sql.DB, arg models.GetUserIdParams) (models.ProfileResponse, error) {
	row := db.QueryRowContext(context.Background(), getUserId, arg.ID, arg.IsActive)
	var i models.ProfileResponse
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
		&i.TwoFactorEnabled,
		&i.IsActive,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username = $1, phone = $2, fullname = $3, avatar = $4, gender = $5, hidden_phone_number = $6
WHERE id = $7
RETURNING id, username, hidden_phone_number, fullname, avatar, gender
`

func UpdateUser(db *sql.DB, arg models.UpdateUserParams) (models.UpdateUserRow, error) {
	// Start with the base update statement
	updateUser := "UPDATE users SET"

	// Use a slice to hold the values to be updated
	var updateValues []interface{}

	// Use a counter for the placeholder values
	counter := 1

	// Check each field and add it to the update statement if it's provided
	if arg.Username.Valid {
		updateUser += fmt.Sprintf(" username = $%d,", counter)
		updateValues = append(updateValues, arg.Username)
		counter++
	}

	if arg.Phone.Valid {
		updateUser += fmt.Sprintf(" phone = $%d,", counter)
		updateValues = append(updateValues, arg.Phone)
		counter++
	}

	if arg.Fullname.Valid {
		updateUser += fmt.Sprintf(" fullname = $%d,", counter)
		updateValues = append(updateValues, arg.Fullname)
		counter++
	}

	if arg.Avatar.Valid {
		updateUser += fmt.Sprintf(" avatar = $%d,", counter)
		updateValues = append(updateValues, arg.Avatar)
		counter++
	}

	if arg.Gender.Valid {
		updateUser += fmt.Sprintf(" gender = $%d,", counter)
		updateValues = append(updateValues, arg.Gender)
		counter++
	}

	// Remove the trailing comma
	updateUser = updateUser[:len(updateUser)-1]

	// Add the WHERE clause
	updateUser += fmt.Sprintf(" WHERE id = $%d RETURNING id, username, hidden_phone_number, fullname, avatar, gender", counter)
	updateValues = append(updateValues, arg.ID)

	// Execute the query
	row := db.QueryRowContext(context.Background(), updateUser, updateValues...)

	var i models.UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HiddenPhoneNumber,
		&i.Fullname,
		&i.Avatar,
		&i.Gender,
	)
	return i, err
}

const updateTwoFactorEnable = `-- name: UpdateTwoFactorEnable :exec
UPDATE users
SET two_factor_enabled = $1
WHERE id = $2
`

func UpdateTwoFactorEnable(db *sql.DB, arg models.UpdateTwoFactorEnableParams) error {
	_, err := db.ExecContext(context.Background(), updateTwoFactorEnable, arg.TwoFactorEnabled, arg.ID)
	return err
}

// CheckEmailExists checks if an email exists in the database.
// It takes a database connection and a CheckEmailExistsParams struct as arguments.
// It returns a boolean indicating whether the email exists and an error, if any.
const checkEmailExists = `-- name: CheckEmailExists :one
SELECT EXISTS (
    SELECT 1 
    FROM users 
    WHERE email = $1 
    AND id != $2
) AS email_exists
`

func CheckEmailExists(db *sql.DB, arg models.CheckEmailExistsParams) (bool, error) {
	row := db.QueryRowContext(context.Background(), checkEmailExists, arg.Email, arg.ID)
	var email_exists bool
	err := row.Scan(&email_exists)
	return email_exists, err
}

const updateEmail = `-- name: UpdateEmail :exec
UPDATE users
SET email = $1, hidden_email = $2
WHERE id = $3
`

func UpdateEmail(db *sql.DB, arg models.UpdateEmailParams) error {
	_, err := db.ExecContext(context.Background(), updateEmail, arg.Email, arg.HiddenEmail, arg.ID)
	return err
}
