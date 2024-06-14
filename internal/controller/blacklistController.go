package controller

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/service"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// BlackListIP adds the client's IP address to the blacklist.
// It calls the BlackListIP function from the service package to perform the operation.
// If the operation is successful, it returns a success response with the result.
// If the operation fails, it returns an error.
func BlackListIP(c *gin.Context) error {
	result := service.BlackListIP(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Add Black List", result)
	return nil
}
