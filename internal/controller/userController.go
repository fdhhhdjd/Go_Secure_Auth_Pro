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

// UpdateProfile updates the profile of a user.
// It retrieves the user's profile using the service.GetProfileUser function,
// and if the result is not nil, it sends a successful response with the updated profile.
// If the result is nil, it returns nil.
func UpdateProfile(c *gin.Context) error {
	result := service.UpdateProfileUser(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Update Profile User", result)
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

// ChangePassword is a controller function that handles the change password request.
// It calls the service.ChangePassword function to perform the password change operation.
// If the operation is successful, it returns a success response with the updated user information.
// If the operation fails, it returns an error response.
func ChangePassword(c *gin.Context) error {
	result := service.ChangePassword(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Change Password User", result)
	return nil
}

// EnableTowFactor enables two-factor authentication for a user.
// It calls the ChangePassword function from the service package to change the user's password.
// If the password change is successful, it returns a success response with the updated user information.
// Otherwise, it returns an error.
func EnableTowFactor(c *gin.Context) error {
	result := service.EnableTowFactor(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Enable Tow Factor User", result)
	return nil
}

// SendOtpUpdateEmail sends an OTP (One-Time Password) to update the user's email.
// It calls the SendOtpUpdateEmail function from the service package and returns the result.
// If the result is nil, it returns nil. Otherwise, it sends a success response with the result.
func SendOtpUpdateEmail(c *gin.Context) error {
	result := service.SendOtpUpdateEmail(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Send Otp Update Email", result)
	return nil
}

// UpdateEmailUser updates the email of a user.
// It calls the service.UpdateEmailUser function to perform the update operation.
// If the update is successful, it returns a success response with the updated user information.
// If there is an error during the update, it returns the error.
func UpdateEmailUser(c *gin.Context) error {
	result := service.UpdateEmailUser(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Update Email User", result)
	return nil
}
