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
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	Token          string    `json:"token"`
	ExpiresAtToken time.Time `json:"expires_at_token"`
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

type LoginTwoFactor struct {
	ID        int       `json:"id"`
	DeviceID  string    `json:"device_id"`
	Email     string    `json:"email"`
	Code      int       `json:"code"`
	ExpiredAt time.Time `json:"expired_at"`
}

// * ---Login Social
type BodyLoginSocialRequest struct {
	Uid  string `json:"uid" binding:"required"`
	Type int    `json:"type" binding:"required"`
}

type SocialResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
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

type UpdateUserRow struct {
	ID                int32          `json:"id"`
	Username          sql.NullString `json:"username"`
	HiddenPhoneNumber sql.NullString `json:"hidden_phone_number"`
	Fullname          sql.NullString `json:"fullname"`
	Avatar            sql.NullString `json:"avatar"`
	Gender            sql.NullInt32  `json:"gender"`
}

type UpdateUserParams struct {
	Username          sql.NullString `json:"username"`
	Phone             sql.NullString `json:"phone"`
	Fullname          sql.NullString `json:"fullname"`
	Avatar            sql.NullString `json:"avatar"`
	Gender            sql.NullInt64  `json:"gender"`
	HiddenPhoneNumber sql.NullString `json:"hidden_phone_number"`
	ID                int            `json:"id"`
}

type BodyUpdateRequest struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	FullName string `json:"fullname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
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

// * --- Forget Password
type ForgetResponse struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type BodyForgetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// * --- Reset Password
type BodyResetPasswordRequest struct {
	Token    string `json:"token" binding:"required"`
	UserId   int    `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type ResetPasswordResponse struct {
	Id int `json:"id"`
}

type UpdateOnlyPasswordParams struct {
	PasswordHash string `json:"password_hash"`
	ID           int    `json:"id"`
}

// * --- Profile
type ProfileResponse struct {
	ID                int            `json:"id"`
	Username          sql.NullString `json:"username"`
	Email             string         `json:"email"`
	Phone             sql.NullString `json:"phone"`
	HiddenPhoneNumber sql.NullString `json:"hidden_phone_number"`
	FullName          sql.NullString `json:"fullname"`
	HiddenEmail       sql.NullString `json:"hidden_email"`
	Avatar            sql.NullString `json:"avatar"`
	Gender            sql.NullInt16  `json:"gender"`
	TwoFactorEnabled  bool           `json:"two_factor_enabled"`
	IsActive          bool           `json:"is_active"`
	CreatedAt         time.Time      `json:"created_at"`
}

type ProfileResponseJSON struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	HiddenPhoneNumber string `json:"hidden_phone_number"`
	FullName          string `json:"fullname"`
	HiddenEmail       string `json:"hidden_email"`
	Avatar            string `json:"avatar"`
	Gender            int    `json:"gender"`
	TwoFactorEnabled  bool   `json:"two_factor_enabled"`
	IsActive          bool   `json:"is_active"`
	CreatedAt         string `json:"created_at"`
}

type PramsProfileRequest struct {
	Id int `uri:"id" binding:"required,min=1"`
}

type GetUserIdParams struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active"`
}

// * --- Logout
type LogoutResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

//* --- Renew Token

type PayloadRefetchResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

//* --- Change Password

type ChangePassResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type BodyChangePasswordRequest struct {
	Password string `json:"password" binding:"required,min=6"`
}

// * Update Two Factor Enable
type UpdateTwoFactorEnableParams struct {
	TwoFactorEnabled bool `json:"two_factor_enabled"`
	ID               int  `json:"id"`
}

type BodyTwoFactorEnableRequest struct {
	TwoFactorEnabled bool `json:"two_factor_enabled"`
}

// * Update Email
type UpdateEmailParams struct {
	Email       string `json:"email" binding:"required,email"`
	HiddenEmail string `json:"hidden_email"`
	ID          int    `json:"id"`
}

type CheckEmailExistsParams struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
}

type BodyUpdateEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
	Otp   string `json:"otp" binding:"required"`
}

type DestroyAccountResponse struct {
	Id int `json:"id"`
}
