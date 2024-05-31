package constants

import "time"

const (
	DevEnvironment = "dev"
)

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
	NotNullViolation    = "23502"
)

const (
	ExitData    = 1
	NotExitData = 0
)

const (
	Verification   = 10
	ResetPassword  = 20
	ChangePassword = 30
)

const (
	Email    = 1
	Phone    = 2
	Username = 3
)

const (
	KeyRandomTokenVerification = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

const (
	ExpiresAccessToken  = 15 * time.Minute
	ExpiresRefreshToken = 7 * 24 * time.Hour
)
