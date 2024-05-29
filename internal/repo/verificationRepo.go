package repo

import (
	"context"
	"database/sql"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
)

// CreateVerification creates a new verification record in the database.
// It takes a database connection `db` and a `data` object of type `models.BodyVerificationRequest`
// containing the necessary information for creating the verification record.
// It returns a `models.Verification` object representing the created verification record and an error, if any.
func CreateVerification(db *sql.DB, data models.BodyVerificationRequest) (models.Verification, error) {
	row := db.QueryRow("INSERT INTO verification (user_id, verified_token, expires_at) "+
		"VALUES ($1, $2, $3) RETURNING id", data.UserId, data.VerifiedToken, data.ExpiresAt)
	var i models.Verification
	err := row.Scan(
		&i.ID,
	)
	return i, err
}

// GetVerification retrieves the verification details from the database based on the provided token and user ID.
// It returns a models.Verification object and an error if any.

const getVerification = `-- name: GetVerification :one
SELECT id, user_id, verified_token, is_verified, verified_at, expires_at, is_active, created_at, updated_at FROM verification
WHERE verified_token = $1 AND user_id = $2 AND is_verified = $3 LIMIT 1
`

func GetVerification(db *sql.DB, arg models.QueryLoginRequest) (models.Verification, error) {

	row := db.QueryRowContext(context.Background(), getVerification, arg.Token, arg.UserId, false)
	var i models.Verification
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.VerifiedToken,
		&i.IsVerified,
		&i.VerifiedAt,
		&i.ExpiresAt,
		&i.IsActive,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateVerification = `-- name: UpdateVerification :exec
UPDATE verification
SET is_verified = $1, is_active = $2
WHERE user_id = $3
`

// UpdateVerification updates the verification status and activity status of a user in the database.
// It takes a database connection and the necessary parameters as arguments.
// Returns an error if the database update fails.
func UpdateVerification(db *sql.DB, arg models.UpdateVerificationParams) error {
	_, err := db.ExecContext(context.Background(), updateVerification, arg.IsVerified, arg.IsActive, arg.UserID)
	return err
}
