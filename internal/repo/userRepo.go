package repo

import (
	"database/sql"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
)

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

func CreateUser(db *sql.DB, email string) (models.User, error) {
	row := db.QueryRow("INSERT INTO users (email) VALUES ($1) RETURNING id", email)
	var i models.User
	err := row.Scan(
		&i.ID,
	)
	return i, err
}
