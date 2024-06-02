package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID                int            `json:"id"`
	Username          sql.NullString `json:"username"`
	Email             string         `json:"email"`
	Phone             sql.NullString `json:"phone"`
	HiddenPhoneNumber sql.NullString `json:"hidden_phone_number"`
	FullName          sql.NullString `json:"fullname"`
	HiddenEmail       sql.NullString `json:"hidden_email"`
	Avatar            sql.NullString `json:"avatar"`
	Gender            sql.NullInt16  `json:"gender"`
	PasswordHash      sql.NullString `json:"password_hash"`
	TwoFactorEnabled  bool           `json:"two_factor_enabled"`
	IsActive          bool           `json:"is_active"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

// * --- Register
type BodyRegisterRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type RegistrationResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
type TokenVerificationLink struct {
	Token string
	Link  string
}

// * --- Login
type BodyLoginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	ID          int    `json:"id"`
	DeviceID    string `json:"device_id"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

// * --- UpdateUser
type UpdatePasswordParams struct {
	PasswordHash string `json:"password_hash"`
	ID           int    `json:"id"`
	Username     string `json:"username"`
	FullName     string `json:"fullname"`
	HiddenEmail  string `json:"hidden_email"`
	IsActive     bool   `json:"is_active"`
}

type UpdateUserResponse struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	HiddenEmail string `json:"hidden_email"`
	IsActive    bool   `json:"is_active"`
}

// *  --- Payload Token
type Payload struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type UserIDEmail struct {
	ID    int
	Email string
}

// * --- Spam user Redis
type SpamUserResponse struct {
	ExpiredSpam int  `json:"expired_spam"`
	IsSpam      bool `json:"is_spam"`
}
