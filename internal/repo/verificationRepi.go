package repo

import (
	"database/sql"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
)

func CreateVerification(db *sql.DB, data models.BodyVerificationRequest) (models.Verification, error) {
	row := db.QueryRow("INSERT INTO verification (user_id, verified_token, expires_at) "+
		"VALUES ($1, $2, $3) RETURNING id", data.UserId, data.VerifiedToken, data.ExpiresAt)
	var i models.Verification
	err := row.Scan(
		&i.ID,
	)
	return i, err
}
