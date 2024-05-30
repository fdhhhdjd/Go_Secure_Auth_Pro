package models

import (
	"time"
)

type Verification struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	VerifiedToken string    `json:"verified_token"`
	IsVerified    bool      `json:"is_verified"`
	VerifiedAt    time.Time `json:"verified_at"`
	ExpiresAt     time.Time `json:"expires_at"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type BodyVerificationRequest struct {
	UserId        int       `json:"user_id" binding:"required"`
	VerifiedToken string    `json:"verified_token" binding:"required"`
	ExpiresAt     time.Time `json:"expires_at" binding:"required"`
}

type UpdateVerificationParams struct {
	IsVerified bool `json:"is_verified"`
	IsActive   bool `json:"is_active"`
	UserID     int  `json:"user_id"`
}

type VerificationResponse struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}

type QueryVerificationRequest struct {
	UserId int    `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
	Email  string `form:"email" binding:"required"`
}
