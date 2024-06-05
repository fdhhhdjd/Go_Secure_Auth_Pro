package service

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// LoginSocial handles the login process for social authentication.
// It takes a gin.Context object as a parameter and returns a pointer to a models.LoginResponse object.
// The function first binds the JSON request body to a models.BodyLoginSocialRequest object.
// If there is an error in binding the JSON, it returns a bad request error response and nil.
// Then, based on the social authentication type, it calls the corresponding social authentication function.
// If the social authentication type is not supported, it returns a bad request error response.
// Next, it joins the users table with the verification table using the user's email.
// If there is an error in joining the tables or no users are found, it returns a bad request error response and nil.
// Otherwise, it retrieves the first user from the result and checks if the user's account is active.
// If the account is blocked, it returns a forbidden error response and nil.
// It then creates an access token, a refresh token, and encodes the public key using the user's ID and email.
// If any of the tokens or the encoded public key is empty, it returns a bad request error response and nil.
// Next, it updates the user's device information and returns the device ID.
// Finally, it sets a cookie with the refresh token and returns a models.LoginResponse object with the user's ID, device ID, email, and access token.
func LoginSocial(c *gin.Context) *models.LoginResponse {
	reqBody := models.BodyLoginSocialRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.BadRequestError(c)
		return nil
	}

	var resultInfoSocial *models.SocialResponse

	switch reqBody.Type {
	case constants.SocialGoogle:
		resultInfoSocial = socialGoogle(c, reqBody.Uid)
	default:
		response.BadRequestError(c)
	}

	users, err := repo.JoinUsersWithVerificationByEmail(global.DB, resultInfoSocial.Email)

	if err != nil {
		response.BadRequestError(c)
		return nil
	}

	if len(users) == 0 {
		response.BadRequestError(c)
		return nil
	}

	resultUser := &users[0]

	accountBlock := CheckUserIsActive(resultUser.IsActive)
	if accountBlock == nil {
		response.ForbiddenError(c)
		return nil
	}

	accessToken, refetchToken, resultEncodePublicKey := createKeyAndToken(models.UserIDEmail{
		ID:    resultUser.ID,
		Email: resultUser.Email,
	})

	if accessToken == "" || refetchToken == "" || resultEncodePublicKey == "" {
		response.BadRequestError(c)
		return nil
	}

	resultInfoDevice := upsetDevice(c, resultUser.ID, resultEncodePublicKey)

	setCookie(c, constants.UserLoginKey, refetchToken, "/", constants.AgeCookie)

	return &models.LoginResponse{
		ID:          resultUser.ID,
		DeviceID:    resultInfoDevice.DeviceID,
		Email:       resultUser.Email,
		AccessToken: accessToken,
	}
}

// socialGoogle retrieves social information from Google for a given user ID.
// It takes a Gin context and a user ID as input and returns a SocialResponse object.
// If the user ID token is invalid or missing, it returns a nil response.
func socialGoogle(c *gin.Context, uid string) *models.SocialResponse {
	infoUserSocial := helpers.GetUserIDToken(c, uid)

	if infoUserSocial == nil {
		response.BadRequestError(c)
		return nil
	}

	return &models.SocialResponse{
		Fullname: infoUserSocial.Fullname,
		Email:    infoUserSocial.Email,
		Picture:  infoUserSocial.Picture,
	}
}
