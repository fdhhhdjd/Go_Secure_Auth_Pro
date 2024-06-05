package controller

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/service"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// Register handles the registration process for a user.
// It calls the Register function from the service package and returns the result.
// If the result is nil, it returns nil.
// Otherwise, it sends a "Created" response with the result.
func Register(c *gin.Context) error {
	result := service.Register(c)
	if result == nil {
		return nil
	}
	response.Created(c, "Register", result)
	return nil
}

// ResendVerificationLink resends the verification link to the user.
func ResendVerificationLink(c *gin.Context) error {
	result := service.ResendVerificationLink(c)
	if result == nil {
		return nil
	}
	response.Created(c, "Resend Verification", result)
	return nil
}

// VerificationAccount handles the verification of user accounts.
// It calls the service.VerificationAccount function to perform the verification.
// If the verification is successful, it sends a success response using the response.Ok function.
// If an error occurs during the verification process, it returns the error.
func VerificationAccount(c *gin.Context) error {
	result := service.VerificationAccount(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Verification", result)
	return nil
}

// LoginIdentifier handles the login identifier request.
// It calls the LoginIdentifier function from the service package to perform the login identifier logic.
// If the result is not nil, it sends a successful response with the result.
// Returns an error if there was an issue during the process.
func LoginIdentifier(c *gin.Context) error {
	result := service.LoginIdentifier(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Login Identifier", result)
	return nil
}

// LoginSocial handles the social login functionality.
// It calls the LoginSocial function from the service package to perform the login operation.
// If the login is successful, it returns the result as a response using the Ok function from the response package.
// If the login fails, it returns an error.
func LoginSocial(c *gin.Context) error {
	result := service.LoginSocial(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Login Social", result)
	return nil
}

// ForgetPassword handles the forget password functionality.
// It calls the service to process the forget password request and returns the result.
func ForgetPassword(c *gin.Context) error {
	result := service.ForgetPassword(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Forget Password", result)
	return nil
}

// ResetPassword handles the reset password functionality.
// It calls the service to reset the password and returns the result.
func ResetPassword(c *gin.Context) error {
	result := service.ResetPassword(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Reset Password", result)
	return nil
}

// RenewToken renews the authentication token.
// It calls the RenewToken function from the service package to renew the token.
// If the token renewal is successful, it sends an "OK" response with the renewed token.
// If there is an error during the token renewal process, it returns the error.
func RenewToken(c *gin.Context) error {
	result := service.RenewToken(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Renew Token", result)
	return nil
}
