package middlewares

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// RefetchTokenMiddleware is a middleware function that handles the refetch token logic.
// It checks if the refetch token is valid and associated with the correct device and user.
// If the refetch token is invalid or the device/user is unauthorized, it aborts the request with an unauthorized status.
// Otherwise, it sets the refetch token in the context and proceeds to the next middleware or handler.
func RefetchTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		refetchToken, err := c.Cookie("user_login")
		if err != nil {
			response.UnauthorizedError(c, response.ErrCookieInvalid)
			return
		}

		deviceID, exists := c.Get("device_id")

		if deviceID == nil || !exists {
			response.UnauthorizedError(c, response.ErrCodeHeaderNotExit)
			return
		}

		resultDevice, errDevice := repo.GetDeviceId(global.DB, models.GetDeviceIdParams{
			DeviceId: deviceID.(string),
			IsActive: true,
		})

		if errDevice != nil {
			response.UnauthorizedError(c, response.ErrCodeDeviceNotExit)
			return
		}

		DecodePublicKeyFromPem, _ := helpers.DecodePublicKeyFromPem(resultDevice.PublicKey.String)

		payload, err := helpers.VerifyToken(refetchToken, DecodePublicKeyFromPem)
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
		userId := userInfo["id"].(float64)
		email := userInfo["email"].(string)

		resultCheckUser := CheckUser(email)

		if !resultCheckUser {
			response.UnauthorizedError(c, response.ErrUserNotExit)
			return
		}

		if int(userId) != resultDevice.UserID {
			response.UnauthorizedError(c, response.ErrUserNotExit)
			return
		}

		c.Set(constants.InfoRefetch, models.PayloadRefetchResponse{
			ID:    int(userId),
			Email: email,
		})

		c.Next()
	}
}
