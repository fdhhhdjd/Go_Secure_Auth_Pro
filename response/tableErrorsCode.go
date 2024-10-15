package response

const (
	//* User Table Errors
	// ErrUserNotExit indicates the account not exit in db
	ErrUserNotExit = 12000

	// ErrUserExit indicates the account already exists in the users table.
	UserExit = 12013

	//ErrUserDuplicateEmail indicates the email already exists in the users table.
	ErrUserDuplicateEmail = 12001

	// ErrUserNotExitEmail indicates the email not exit in db
	ErrUserNotExitEmail = 12002

	// ErrorUserNotExitUsername indicates the username not exit in db
	ErrorUserNotExitUsername = 12003

	// ErrorUserPhoneNotExit indicates the phone not exit in db
	ErrorUserPhoneNotExit = 12004

	// ErrUserNotActive indicates the account not active
	ErrUserNotActive = 12005

	// ErrUserTwoFactorDisabled indicates the two factor disabled
	ErrTwoFactorDisabled = 12006

	// ErrUserTwoFactorInvalid indicates the two factor invalid
	ErrTwoFactorInvalid = 12007

	// ErrUserTwoFactorUnauthorized indicates the two factor unauthorized
	ErrTwoFactorUnauthorized = 12008

	// ErrorUsernameInvalid indicates the username is invalid
	ErrorUsernameInvalid = 12009

	// ErrorUserPhoneInvalid indicates the phone is invalid
	ErrorUserPhoneInvalid = 12010

	// ErrorUserEmailInvalid indicates the email is invalid
	ErrorUserEmailInvalid = 12011

	// ErrTwoFactorEnabled indicates the two factor enabled
	ErrTwoFactorEnabled = 12012

	//* Device Table Errors
	// ErrCodeDeviceNotExit indicates the device not exits
	ErrCodeDeviceNotExit = 12002

	//* Verification Table Errors
	// ErrorVerificationCodeNotExit indicates the verification code not exits
	ErrorVerificationCodeNotExit = 14000

	// ErrorVerificationCodeExpired indicates the verification code has expired
	ErrorVerificationCodeExpired = 14001

	// ErrorVerificationCodeInvalid indicates the verification code is invalid
	ErrorVerificationCodeInvalid = 14002

	// ErrorVerificationCodeDuplicate indicates the verification code is duplicate
	ErrorVerificationCodeDuplicate = 14003

	//* Password Table Errors
	// ErrorPasswordNotExit indicates the password not exits
	ErrorPasswordNotExit = 15000

	// ErrorPasswordNotMatch indicates the password not match
	ErrorPasswordNotMatch = 15001

	// ErrorPasswordNotUpdate indicates the password not update
	ErrorPasswordNotUpdate = 15002

	// ErrorEncryptPassword indicates the password not encrypt
	ErrorEncryptPassword = 15003

	// ErrorPassWeak indicates the password weak
	ErrorPassWeak = 15004

	// ErrorPasswordIsOld indicates the password is old
	ErrorPasswordIsOld = 15005

	//* OTP Table Errors
	// ErrorOTPNotExit indicates the otp not exits
	ErrorOTPNotExit = 16000

	// ErrorOTPExpired indicates the otp has expired
	ErrorOTPExpired = 16001

	// ErrorOTPInvalid indicates the otp is invalid
	ErrorOTPInvalid = 16002
)
