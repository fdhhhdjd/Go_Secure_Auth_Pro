package service

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// GetProfileUser retrieves the profile of a user based on the provided user ID.
// It returns a pointer to a models.ProfileResponse struct.
func GetProfileUser(c *gin.Context) *models.ProfileResponse {
	var req models.PramsProfileRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	resultUser, err := repo.GetUserId(global.DB, models.GetUserIdParams{
		ID:       req.UserId,
		IsActive: true,
	})

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}
	return &resultUser
}

// Logout logs out the user and clears the session.
// It takes a gin.Context as input and returns a pointer to models.LogoutResponse.
// If the user_info or device_id is missing in the context, it returns a BadRequestError.
// Otherwise, it updates the logout time in the database, clears the user login cookie,
// and returns a pointer to models.LogoutResponse containing the user ID and email.
func Logout(c *gin.Context) *models.LogoutResponse {
	payload, existsUserInfo := c.Get(constants.InfoAccess)
	deviceId, existsDevice := c.Get("device_id")

	if !existsUserInfo || !existsDevice {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	repo.UpdateTimeLogout(global.DB, models.UpdateTimeLogoutParams{
		LoggedOutAt: sql.NullTime{Time: time.Now()},
		DeviceId:    deviceId.(string),
	})

	clearCookie(c, constants.UserLoginKey)

	return &models.LogoutResponse{
		Id:    payload.(models.Payload).ID,
		Email: payload.(models.Payload).Email,
	}
}

// clearCookie clears the specified cookie from the response.
// The cookie is set to expire immediately and its value is emptied.
// The `httpOnly` flag is set based on the environment.
func clearCookie(c *gin.Context, cookieName string) {
	nodeEnv := os.Getenv("ENV")
	var httpOnly bool
	if nodeEnv == constants.DevEnvironment {
		httpOnly = false
	} else {
		httpOnly = true
	}

	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: httpOnly,
	}

	http.SetCookie(c.Writer, cookie)
}
