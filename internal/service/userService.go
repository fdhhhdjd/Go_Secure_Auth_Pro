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
func GetProfileUser(c *gin.Context) *models.ProfileResponseJSON {
	var req models.PramsProfileRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	user, err := repo.GetUserId(global.DB, models.GetUserIdParams{
		ID:       req.Id,
		IsActive: true,
	})

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
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
func UpdateProfileUser(c *gin.Context) *models.UpdateUserRow {
	reqBody := models.BodyUpdateRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	if !validate.ValidateAndRespond(reqBody.Username, validate.IsValidateUser) {
		c.JSON(response.StatusBadRequest, response.BadRequestError(constants.UsernameInvalid))
		return nil
	}

	if !validate.ValidateAndRespond(reqBody.Phone, validate.IsValidatePhone) {
		c.JSON(response.StatusBadRequest, response.BadRequestError(constants.PhoneInvalid))
		return nil
	}

	payload, existsUserInfo := c.Get(constants.InfoAccess)

	if !existsUserInfo {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
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
			c.JSON(response.StatusInternalServerError, response.InternalServerError(errorCreateUser))
			return nil
		}

		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	return &resultUpdateProfile
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
func ChangePassword(c *gin.Context) *models.ChangePassResponse {
	reqBody := models.BodyChangePasswordRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError(constants.PasswordInvalid))
		return nil
	}
	payload, existsUserInfo := c.Get(constants.InfoAccess)

	if !existsUserInfo {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	if !validate.ValidateAndRespond(reqBody.Password, validate.IsValidPassword) {
		c.JSON(response.StatusBadRequest, response.BadRequestError(constants.PasswordWeak))
		return nil
	}

	hashedPassword := checkPasswordOld(reqBody.Password, payload.(models.Payload).ID)

	if hashedPassword == nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError(constants.PasswordHasUsed))
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
