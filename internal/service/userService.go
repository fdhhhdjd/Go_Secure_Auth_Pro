package service

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/utils"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers/validate"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// GetProfileUser retrieves the profile of a user based on the provided user ID.
// It returns a pointer to a models.ProfileResponse struct.
//
// @Summary Get user profile
// @Description Retrieves the profile information of a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param X-Device-Id header string true "Device ID"
// @Success 200 {object} models.ProfileResponseJSON
// @Failure 400 {object} response.ErrorResponse
// @Router /users/{id} [get]
func GetProfileUser(c *gin.Context) *models.ProfileResponseJSON {
	var req models.PramsProfileRequest

	if err := c.ShouldBindUri(&req); err != nil {
		response.BadRequestError(c)
		return nil
	}

	user, err := repo.GetUserId(global.DB, models.GetUserIdParams{
		ID:       req.Id,
		IsActive: true,
	})

	if err != nil {
		response.BadRequestError(c)
		return nil
	}
	response := &models.ProfileResponseJSON{
		ID:                user.ID,
		Username:          helpers.NullStringToString(user.Username),
		Email:             user.Email,
		Phone:             helpers.NullStringToString(user.Phone),
		HiddenPhoneNumber: helpers.NullStringToString(user.HiddenPhoneNumber),
		FullName:          helpers.NullStringToString(user.FullName),
		HiddenEmail:       helpers.NullStringToString(user.HiddenEmail),
		Avatar:            helpers.NullStringToString(user.Avatar),
		Gender:            helpers.NullInt16ToString(user.Gender),
		TwoFactorEnabled:  strconv.FormatBool(user.TwoFactorEnabled),
		IsActive:          strconv.FormatBool(user.IsActive),
		CreatedAt:         user.CreatedAt.Format(time.RFC3339),
	}

	return response
}

// UpdateProfileUser updates the profile of a user based on the provided request body.
// It validates the request body fields, checks the user's access information, and updates the user's profile in the database.
// If any validation or database error occurs, it returns an appropriate error response.
// Otherwise, it returns the updated user profile.
//
// @Summary Update user profile
// @Description Updates the profile information of a user
// @Tags Users
// @Accept json
// @Produce json
// @Param X-Device-Id header string true "Device ID"
// @Param body body models.BodyUpdateRequest true "Profile update request body"
// @Success 200 {object} models.UpdateUserRow
// @Failure 400 {object} response.ErrorResponse
// @Router /users/profile [post]
func UpdateProfileUser(c *gin.Context) *models.UpdateUserRow {
	reqBody := models.BodyUpdateRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.BadRequestError(c)
		return nil
	}

	if !validate.ValidateAndRespond(reqBody.Username, validate.IsValidateUser) {
		response.BadRequestError(c, constants.UsernameInvalid)
		return nil
	}

	if !validate.ValidateAndRespond(reqBody.Phone, validate.IsValidatePhone) {
		response.BadRequestError(c, constants.PhoneInvalid)
		return nil
	}

	payload, existsUserInfo := c.Get(constants.InfoAccess)

	if !existsUserInfo {
		response.BadRequestError(c)
		return nil
	}

	resultUpdateProfile, err := repo.UpdateUser(global.DB, models.UpdateUserParams{
		Username:          sql.NullString{String: reqBody.Username, Valid: reqBody.Username != ""},
		Phone:             sql.NullString{String: reqBody.Phone, Valid: reqBody.Phone != ""},
		Fullname:          sql.NullString{String: reqBody.FullName, Valid: reqBody.FullName != ""},
		Avatar:            sql.NullString{String: reqBody.Avatar, Valid: reqBody.Avatar != ""},
		Gender:            sql.NullInt64{Int64: int64(reqBody.Gender), Valid: true},
		HiddenPhoneNumber: sql.NullString{String: helpers.HidePhoneNumber(reqBody.Phone), Valid: reqBody.Phone != ""},
		ID:                payload.(models.Payload).ID,
	})

	if err != nil {
		//* Error for database
		errorCreateUser := utils.HandleDBError(err)
		if errorCreateUser != "" {
			response.InternalServerError(c, errorCreateUser)
			return nil
		}

		response.BadRequestError(c)
		return nil
	}

	return &resultUpdateProfile
}

// Logout logs out the user and clears the session.
// It takes a gin.Context as input and returns a pointer to models.LogoutResponse.
// If the user_info or device_id is missing in the context, it returns a BadRequestError.
// Otherwise, it updates the logout time in the database, clears the user login cookie,
// and returns a pointer to models.LogoutResponse containing the user ID and email.
//
// @Summary Logout user
// @Description Logs out a user
// @Tags Users
// @Accept json
// @Produce json
// @Param X-Device-Id header string true "Device ID"
// @Success 200 {object} models.LogoutResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /logout [post]
func Logout(c *gin.Context) *models.LogoutResponse {
	payload, existsUserInfo := c.Get(constants.InfoAccess)
	deviceId, existsDevice := c.Get("device_id")

	if !existsUserInfo || !existsDevice {
		response.BadRequestError(c)
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

// ChangePassword is a function that handles the change password request.
// It takes a gin.Context object as a parameter and returns a pointer to a models.ChangePassResponse object.
// The function first binds the JSON request body to a models.BodyChangePasswordRequest object.
// If the binding fails, it returns a BadRequestError response with the error message "PasswordInvalid".
// It then checks if the user information exists in the context.
// If not, it returns a BadRequestError response.
// Next, it validates the password using the validate.IsValidPassword function.
// If the password is weak, it returns a BadRequestError response with the error message "PasswordWeak".
// It then checks the old password against the hashed password stored in the payload.
// If the old password is not valid, it returns a BadRequestError response with the error message "PasswordHasUsed".
// The function inserts the old password into the password history table and updates the password in the user table.
// Finally, it returns a ChangePassResponse object with the user ID and email.
//
// @Summary Change user password
// @Description Changes the password for a user
// @Tags Users
// @Accept json
// @Produce json
// @Param X-Device-Id header string true "Device ID"
// @Param body body models.BodyChangePasswordRequest true "Password change request body"
// @Success 200 {object} models.ChangePassResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /change-password [post]
func ChangePassword(c *gin.Context) *models.ChangePassResponse {
	reqBody := models.BodyChangePasswordRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.BadRequestError(c, constants.PasswordInvalid)
		return nil
	}
	payload, existsUserInfo := c.Get(constants.InfoAccess)

	if !existsUserInfo {
		response.BadRequestError(c)
		return nil
	}

	if !validate.ValidateAndRespond(reqBody.Password, validate.IsValidPassword) {
		response.BadRequestError(c, constants.PasswordWeak)
		return nil
	}

	hashedPassword := checkPasswordOld(reqBody.Password, payload.(models.Payload).ID)

	if hashedPassword == nil {
		response.BadRequestError(c, constants.PasswordHasUsed)
		return nil
	}

	repo.InsertPasswordHistory(global.DB, models.InsertPasswordHistoryParams{
		UserID:       payload.(models.Payload).ID,
		OldPassword:  hashedPassword.Salt,
		ReasonStatus: constants.ResetPassword,
	})

	repo.UpdateOnlyPassword(global.DB, models.UpdateOnlyPasswordParams{
		ID:           payload.(models.Payload).ID,
		PasswordHash: hashedPassword.HashedPassword,
	})

	return &models.ChangePassResponse{
		Id:    payload.(models.Payload).ID,
		Email: payload.(models.Payload).Email,
	}
}

// EnableTowFactor enables two-factor authentication for a user.
// It takes a Gin context `c` and returns an `UpdateTwoFactorEnableParams` pointer.
// The function first parses the JSON request body into a `BodyTwoFactorEnableRequest` struct.
// If the JSON parsing fails, it responds with a bad request error and returns nil.
// Then, it retrieves the user information from the Gin context.
// If the user information does not exist, it responds with a bad request error and returns nil.
// Finally, it updates the two-factor authentication status for the user in the database
// and returns an `UpdateTwoFactorEnableParams` pointer with the updated information.
//
// @Summary Enable two-factor authentication
// @Description Enables two-factor authentication for a user
// @Tags Users
// @Accept json
// @Produce json
// @Param X-Device-Id header string true "Device ID"
// @Param body body models.BodyTwoFactorEnableRequest true "Two-factor enable request body"
// @SecurityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
// @Success 200 {object} models.UpdateTwoFactorEnableParams
// @Failure 400 {object} response.ErrorResponse
// @Router /enable-two-factor [post]
func EnableTowFactor(c *gin.Context) *models.UpdateTwoFactorEnableParams {
	reqBody := models.BodyTwoFactorEnableRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.BadRequestError(c, constants.TwoFactorEnabledInvalid)
		return nil
	}

	payload, existsUserInfo := c.Get(constants.InfoAccess)

	if !existsUserInfo {
		response.BadRequestError(c)
		return nil
	}

	repo.UpdateTwoFactorEnable(global.DB, models.UpdateTwoFactorEnableParams{
		ID:               payload.(models.Payload).ID,
		TwoFactorEnabled: reqBody.TwoFactorEnabled,
	})
	return &models.UpdateTwoFactorEnableParams{
		ID:               payload.(models.Payload).ID,
		TwoFactorEnabled: reqBody.TwoFactorEnabled,
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
