package middlewares

import (
	"strings"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthorizationMiddleware is a middleware function that handles authorization logic.
// It checks the Authorization header and device ID in the request context to ensure the request is authorized.
// If the request is not authorized, it aborts the request with a JSON response containing an unauthorized error.
// It also verifies the access token and checks if the user email and ID match the device information.
// If all checks pass, it sets the user information in the request context and proceeds to the next middleware or handler.
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		refetchToken, _ := c.Cookie("user_login")

		deviceID, exists := c.Get("device_id")

		if authHeader == "" || deviceID == nil || !exists || refetchToken == "" {
			response.UnauthorizedError(c, response.ErrCodeAuthTokenInvalid)
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			response.UnauthorizedError(c, response.ErrCodeHeaderNotExit)
			return
		}

		resultDevice, err := repo.GetDeviceId(global.DB, models.GetDeviceIdParams{
			DeviceId: deviceID.(string),
			IsActive: true,
		})

		if err != nil || resultDevice.PublicKey.String == "" {
			response.UnauthorizedError(c, response.ErrCodeAuthTokenInvalid)
			return
		}

		accessToken := fields[1]

		DecodePublicKeyFromPem, _ := helpers.DecodePublicKeyFromPem(resultDevice.PublicKey.String)

		payload, err := helpers.VerifyToken(accessToken, DecodePublicKeyFromPem)
		if err != nil {
			response.UnauthorizedError(c, response.ErrCodeAuthTokenInvalid)
			return
		}

		claims, ok := payload.Claims.(jwt.MapClaims)
		if !ok {
			response.UnauthorizedError(c, response.ErrCodeAuthTokenInvalid)
			return
		}

		userInfo := claims["userInfo"].(map[string]interface{})
		email := userInfo["email"].(string)
		userId := userInfo["id"].(float64)

		resultCheckUser := CheckUser(email)

		if !resultCheckUser {
			response.UnauthorizedError(c, response.ErrUserNotExit)
			return
		}

		if int(userId) != resultDevice.UserID {
			response.UnauthorizedError(c, response.ErrUserNotExit)
			return
		}

		c.Set(constants.InfoAccess, models.Payload{
			ID:    int(userId),
			Email: email,
		})

		c.Next()

	}
}

// checkUser checks if a user is valid and active based on the provided email.
// It retrieves the user details from the repository and returns true if the user is valid and active, false otherwise.
func CheckUser(email string) bool {
	resultDetailUser, err := repo.GetUserDetail(global.DB, email)

	if err != nil {
		return false
	}

	if !resultDetailUser.IsActive {
		return false
	}

	return true
}
