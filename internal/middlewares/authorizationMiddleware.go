package middlewares

import (
	"strings"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		deviceID, exists := c.Get("device_id")

		if authHeader == "" || deviceID == nil || !exists {
			c.AbortWithStatusJSON(response.StatusUnauthorized, response.UnauthorizedError())
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(response.StatusUnauthorized, response.UnauthorizedError())
			return
		}

		resultDevice, err := repo.GetDeviceId(global.DB, models.GetDeviceIdParams{
			DeviceId: deviceID.(string),
			IsActive: true,
		})

		if err != nil || resultDevice.PublicKey.String == "" {
			c.AbortWithStatusJSON(response.StatusUnauthorized, response.UnauthorizedError())
			return
		}

		accessToken := fields[1]

		DecodePublicKeyFromPem, _ := helpers.DecodePublicKeyFromPem(resultDevice.PublicKey.String)

		payload, err := helpers.VerifyToken(accessToken, DecodePublicKeyFromPem)
		if err != nil {
			c.AbortWithStatusJSON(response.StatusUnauthorized, response.UnauthorizedError())
			return
		}

		claims, ok := payload.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(response.StatusUnauthorized, response.UnauthorizedError())
			return
		}

		userInfo := claims["userInfo"].(map[string]interface{})
		email := userInfo["email"].(string)
		userId := userInfo["id"].(float64)

		if int(userId) != resultDevice.UserID {
			c.AbortWithStatusJSON(response.StatusUnauthorized, response.UnauthorizedError())
			return
		}

		c.Set("user_info", models.Payload{
			ID:    int(userId),
			Email: email,
		})

		c.Next()

	}
}
