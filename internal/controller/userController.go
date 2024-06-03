package controller

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/service"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// GetProfileUser retrieves the profile of a user.
// It calls the GetProfileUser function from the service package to fetch the user's profile.
// If the result is nil, it returns nil.
// Otherwise, it sends a successful response with the user's profile data.
func GetProfileUser(c *gin.Context) error {
	result := service.GetProfileUser(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Get Profile User", result)
	return nil
}

// LogoutUser handles the logout functionality for a user.
// It calls the Logout function from the service package to perform the logout operation.
// If the logout is successful, it returns a success response.
// Otherwise, it returns an error response.
func LogoutUser(c *gin.Context) error {
	result := service.Logout(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Logout User", result)
	return nil
}
