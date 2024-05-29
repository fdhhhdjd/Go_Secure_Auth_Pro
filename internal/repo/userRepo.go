package repo

import (
	"context"
	"database/sql"

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
SET password_hash = $1
WHERE id = $2
RETURNING id, email
`

func UpdatePassword(db *sql.DB, arg models.UpdatePasswordParams) (models.UpdateUserResponse, error) {
	var i models.UpdateUserResponse
	err := db.QueryRowContext(context.Background(), updatePassword, arg.PasswordHash, arg.ID).Scan(&i.Id, &i.Email)
	return i, err
}
