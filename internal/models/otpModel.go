package models

import (
	"database/sql"
	"time"
)

type Otp struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id"`
	OtpCode   string       `json:"otp_code"`
	CreatedAt sql.NullTime `json:"created_at"`
	IsActive  bool         `json:"is_active"`
	ExpiresAt time.Time    `json:"expires_at"`
}

type SendOtpResponse struct {
	Id        int    `json:"id"`
	Code      string `json:"code"`
	ExpiredAt string `json:"expired_at"`
}

type CreateOtpParams struct {
	UserID    int       `json:"user_id"`
	OtpCode   string    `json:"otp_code"`
	ExpiresAt time.Time `json:"expires_at"`
}

type OtpRequest struct {
	Otp string `json:"otp" binding:"required"`
}

type UpdateOtpIsActiveParams struct {
	IsActive bool   `json:"is_active"`
	OtpCode  string `json:"otp_code"`
}

type GetNewOtpsRow struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id"`
	OtpCode   string       `json:"otp_code"`
	CreatedAt sql.NullTime `json:"created_at"`
	IsActive  bool         `json:"is_active"`
	ExpiresAt time.Time    `json:"expires_at"`
	Email     string       `json:"email"`
}
