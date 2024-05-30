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
