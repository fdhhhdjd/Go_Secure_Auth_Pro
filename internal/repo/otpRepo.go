package repo

import (
	"context"
	"database/sql"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
)

const createOtp = `-- name: CreateOtp :one
INSERT INTO otps (
    user_id,
    otp_code,
    expires_at
) VALUES (
    $1,
    $2,
    $3
) RETURNING id, user_id, otp_code, created_at, expires_at
`

// CreateOtp creates a new OTP (One-Time Password) record in the database.
// It takes a database connection `db` and the OTP parameters `arg` as input.
// It returns the created OTP record and an error (if any).
func CreateOtp(db *sql.DB, arg models.CreateOtpParams) (models.Otp, error) {
	row := db.QueryRowContext(context.Background(), createOtp, arg.UserID, arg.OtpCode, arg.ExpiresAt)
	var i models.Otp
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.OtpCode,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const getNewOtps = `-- name: GetNewOtps :many
SELECT otps.id, otps.user_id, otps.otp_code, otps.created_at, otps.is_active, otps.expires_at, users.email
FROM otps
JOIN users ON otps.user_id = users.id
WHERE otps.expires_at > NOW()
    AND otps.otp_code = $1
    AND otps.is_active = TRUE
`

// GetNewOtps retrieves a list of new OTPs from the database based on the provided OTP code.
// It takes a database connection (`db`) and an OTP code (`otpCode`) as parameters.
// It returns a slice of `models.GetNewOtpsRow` and an error, if any.
func GetNewOtps(db *sql.DB, otpCode string) ([]models.GetNewOtpsRow, error) {
	rows, err := db.QueryContext(context.Background(), getNewOtps, otpCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.GetNewOtpsRow{}
	for rows.Next() {
		var i models.GetNewOtpsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.OtpCode,
			&i.CreatedAt,
			&i.IsActive,
			&i.ExpiresAt,
			&i.Email,
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

const updateOtpIsActive = `-- name: UpdateOtpIsActive :exec
UPDATE otps
SET is_active = $1
WHERE otp_code = $2
`

// UpdateOtpIsActive updates the isActive status of an OTP code in the database.
// It takes a database connection `db` and an `arg` parameter of type `models.UpdateOtpIsActiveParams`.
// It executes a SQL query to update the isActive status of the OTP code in the database.
// Returns an error if the database query fails.
func UpdateOtpIsActive(db *sql.DB, arg models.UpdateOtpIsActiveParams) error {
	_, err := db.ExecContext(context.Background(), updateOtpIsActive, arg.IsActive, arg.OtpCode)
	return err
}
