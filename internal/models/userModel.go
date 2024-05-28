package models

import "time"

type User struct {
	ID                int       `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	HiddenPhoneNumber string    `json:"hidden_phone_number"`
	FullName          string    `json:"fullname"`
	HiddenEmail       string    `json:"hidden_email"`
	Avatar            string    `json:"avatar"`
	Gender            int16     `json:"gender"`
	PasswordHash      string    `json:"password_hash"`
	TwoFactorEnabled  bool      `json:"two_factor_enabled"`
	IsActive          bool      `json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type BodyRegisterRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type BodyLoginRequest struct {
	Identifier string `json:"Identifier" binding:"required,oneof=email phone username"`
	Password   string `json:"password" binding:"required,min=6,regexp=^(?=.*[0-9])(?=.*[A-Z])(?=.*[!@#$%^&*]).*$"`
}

type RegistrationResponse struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
}
