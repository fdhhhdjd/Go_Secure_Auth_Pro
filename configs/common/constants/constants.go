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
	KeyRandomTokenVerification = "0123456789ABC2222D333nguyenTIENTAIabcdefghijklmnopqrstuvwxyz"
)

const (
	ExpiresAccessToken  = 15 * time.Minute
	ExpiresRefreshToken = 7 * 24 * time.Hour
)

const (
	AgeCookie     = 7 * 24 * 60 * 60
	SecondsInADay = "86400"
)

const (
	DeviceId = "X-Device-Id"
)

const (
	UserLoginKey = "user_login"
	InfoAccess   = "info_access"
	InfoRefetch  = "info_refetch"
	CSRFToken    = "CSRF-Token"
)

const (
	StatusForget   = 10
	StatusRegister = 20
	StatusResend   = 30
)

const (
	MB_1 = 1024 * 1024
)
