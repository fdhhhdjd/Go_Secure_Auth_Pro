package controller

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/service"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// RenewToken renews the authentication token.
// It calls the RenewToken function from the service package to renew the token.
// If the token renewal is successful, it sends an "OK" response with the renewed token.
// If there is an error during the token renewal process, it returns the error.
func VerificationOtp(c *gin.Context) error {
	result := service.VerificationOtp(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Two factor login", result)
	return nil
}
