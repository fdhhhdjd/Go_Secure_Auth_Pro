package service

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/utils"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	pkg "github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/mail"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register handles the registration process for a user.
// It receives a Gin context object and returns a RegistrationResponse pointer.
// The function first retrieves the request body data and checks its validity.
// If the body is invalid, it returns a BadRequestError response.
// Next, it checks if the user already exists in the database.
// If the user exists, it returns a BadRequestError response.
// If the user does not exist, it creates a new user in the database.
// If there is an error during user creation, it returns an InternalServerError response.
// After creating the user, it generates a random verification token and sets its expiration time.
// It then links the token with the user and saves the verification details in the database.
// If there is an error during verification creation, it returns an InternalServerError response.
// Finally, it sends an email to the user with the verification link.
// The function returns a RegistrationResponse object containing the user ID, email, and token.
func Register(c *gin.Context) *models.RegistrationResponse {
	//* Get data for body
	reqBody := models.BodyRegisterRequest{}

	//* Check body valid
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	//* Get detail users
	resultDetailUser, err := repo.GetUserDetail(global.DB, reqBody.Email)

	// * Check account exit into yet
	if err != nil {
		errorDetailUser := utils.HandleDBError(err)
		//* Error for database
		if errorDetailUser != "" {
			c.JSON(response.StatusInternalServerError, response.InternalServerError(errorDetailUser))
			return nil
		}
	}

	//* Check user have exit to yet
	if resultDetailUser.ID != constants.NotExitData {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	//* If user not exit create user
	resultCreateUser, err := repo.CreateUser(global.DB, reqBody.Email)
	if err != nil {
		//* Error for database
		errorCreateUser := utils.HandleDBError(err)
		if errorCreateUser != "" {
			c.JSON(response.StatusInternalServerError, response.InternalServerError(errorCreateUser))
			return nil
		}
		c.JSON(response.StatusBadRequest, response.InternalServerError())
		return nil
	}

	resultVerificationLink := createTokenVerificationLink(c, models.UserIDEmail{
		ID:    resultCreateUser.ID,
		Email: reqBody.Email,
	})

	//* Send email
	data := models.EmailData{
		Title:    "Register User!",
		Body:     resultVerificationLink.Link,
		Template: `<h1>{{.Title}}</h1>Verification account: <a href="{{.Body}}">Click here to verify your account</a> </br> <img src="cid:logo" alt="Image" height="200" />`,
	}

	upsetDevice(c, resultCreateUser.ID, "")

	go pkg.SendGoEmail(reqBody.Email, data)

	return &models.RegistrationResponse{
		ID:    resultCreateUser.ID,
		Email: reqBody.Email,
		Token: resultVerificationLink.Token,
	}
}

// VerificationAccount is a function that handles the verification of a user's account.
// It takes a gin.Context object as a parameter and returns a pointer to a models.LoginResponse object.
// The function first binds the query parameters from the request to a models.QueryVerificationRequest object.
// If there is an error in binding the query parameters, it returns a BadRequestError response.
// It then calls the GetVerification function from the repo package to retrieve the verification details.
// If there is an error in retrieving the verification details or if the retrieved details do not match the request parameters,
// it returns a BadRequestError response.
// If the verification token has expired, it returns an UnauthorizedError response.
// It generates a random password, hashes it, and inserts the old password into the password history.
// It updates the verification status of the user to true and deactivates the verification token.
// It updates the user's password with the new hashed password and hidden email.
// It creates an access token, refetch token, and encodes the public key.
// It updates the user's device information and sets a cookie with the refetch token.
// It sends an email to the user with the new password.
// Finally, it returns a LoginResponse object with the user's ID, device ID, email, and access token.
func VerificationAccount(c *gin.Context) *models.LoginResponse {
	reqQuery := models.QueryVerificationRequest{}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	GetVerification, err := repo.GetVerification(global.DB, models.QueryVerificationRequest{
		UserId: reqQuery.UserId,
		Token:  reqQuery.Token,
	})

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	if GetVerification.UserID != reqQuery.UserId || GetVerification.VerifiedToken != reqQuery.Token || !GetVerification.IsActive {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	if GetVerification.ExpiresAt.Unix() < time.Now().Unix() {
		c.JSON(response.StatusBadRequest, response.UnauthorizedError())
		return nil
	}

	randomPassword := helpers.GenerateRandomPassword(10)

	salt, hashedPassword, err := helpers.HashPassword(randomPassword, bcrypt.DefaultCost)

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	errInsertHistoryPassword := repo.InsertPasswordHistory(global.DB, models.InsertPasswordHistoryParams{
		UserID:       reqQuery.UserId,
		OldPassword:  salt,
		ReasonStatus: constants.Verification,
	})

	if errInsertHistoryPassword != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	err = repo.UpdateVerification(global.DB, models.UpdateVerificationParams{
		UserID:     reqQuery.UserId,
		IsVerified: true,
		IsActive:   false,
	})

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	resultUpdateUser, errUpdatePassword := repo.UpdatePassword(global.DB, models.UpdatePasswordParams{
		ID:           reqQuery.UserId,
		PasswordHash: hashedPassword,
		HiddenEmail:  helpers.HideEmail(reqQuery.Email),
		IsActive:     true,
	})

	if errUpdatePassword != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	accessToken, refetchToken, resultEncodePublicKey := createKeyAndToken(models.UserIDEmail{
		ID:    resultUpdateUser.Id,
		Email: resultUpdateUser.Email,
	})

	if accessToken == "" || refetchToken == "" || resultEncodePublicKey == "" {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	resultInfoDevice := upsetDevice(c, resultUpdateUser.Id, resultEncodePublicKey)

	setCookie(c, constants.UserLoginKey, refetchToken, "/cookie", constants.AgeCookie)

	//* Send email
	data := models.EmailData{
		Title:    "Verification Account Success!",
		Body:     randomPassword,
		Template: `<h1>{{.Title}}</h1> <p style="font-size: large;">Thank You, You have verification account success 😊. <br/> Password New: <b>{{.Body}}</b></p>`,
	}

	go pkg.SendGoEmail(resultUpdateUser.Email, data)

	return &models.LoginResponse{
		ID:          resultUpdateUser.Id,
		DeviceID:    resultInfoDevice.DeviceID,
		Email:       resultUpdateUser.Email,
		AccessToken: accessToken,
	}
}

// LoginIdentifier handles the login process for identifying the user based on the provided identifier (email, phone, or username).
// It retrieves the user information from the database based on the identifier and verifies the password.
// If the login is successful, it generates access and refresh tokens, updates the device information, and sets a cookie with the refresh token.
// Finally, it returns a LoginResponse containing the user ID, device ID, email, and access token.
func LoginIdentifier(c *gin.Context) *models.LoginResponse {
	//* Get data for body
	reqBody := models.BodyLoginRequest{}

	//* Check body valid
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}
	var resultUser *models.User

	switch helpers.IdentifyType(reqBody.Identifier) {
	case constants.Email:
		users, err := repo.JoinUsersWithVerificationByEmail(global.DB, reqBody.Identifier)
		if err != nil {
			c.JSON(response.StatusBadRequest, response.BadRequestError())
			return nil
		}
		if len(users) == 0 {
			c.JSON(response.StatusBadRequest, response.BadRequestError())
			return nil
		}
		resultUser = &users[0]
	case constants.Phone:
		users, err := repo.JoinUsersWithVerificationByPhone(global.DB, reqBody.Identifier)
		if err != nil {
			c.JSON(response.StatusBadRequest, response.BadRequestError())
			return nil
		}
		if len(users) == 0 {
			c.JSON(response.StatusBadRequest, response.BadRequestError())
			return nil
		}
		resultUser = &users[0]
	case constants.Username:
		users, err := repo.JoinUsersWithVerificationByUsername(global.DB, reqBody.Identifier)
		if err != nil {
			c.JSON(response.StatusBadRequest, response.BadRequestError())
			return nil
		}
		if len(users) == 0 {
			c.JSON(response.StatusBadRequest, response.BadRequestError())
			return nil
		}
		resultUser = &users[0]
	default:
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	//* Check account have been block
	accountBlock := checkUserIsActive(resultUser.IsActive)
	if accountBlock == nil {
		c.JSON(response.StatusForbidden, response.ForbiddenError())
		return nil
	}

	errPassword := helpers.ComparePassword(reqBody.Password, resultUser.PasswordHash)

	if errPassword != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	accessToken, refetchToken, resultEncodePublicKey := createKeyAndToken(models.UserIDEmail{
		ID:    resultUser.ID,
		Email: resultUser.Email,
	})

	if accessToken == "" || refetchToken == "" || resultEncodePublicKey == "" {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	resultInfoDevice := upsetDevice(c, resultUser.ID, resultEncodePublicKey)

	setCookie(c, constants.UserLoginKey, refetchToken, "/cookie", constants.AgeCookie)

	return &models.LoginResponse{
		ID:          resultUser.ID,
		DeviceID:    resultInfoDevice.DeviceID,
		Email:       resultUser.Email,
		AccessToken: accessToken,
	}
}

// ResendVerificationLink is a function that handles the resend verification link feature.
// It takes a gin.Context object as a parameter and returns a pointer to a models.RegistrationResponse object.
// The function first retrieves the request body data and checks its validity.
// If the request body is invalid, it returns a JSON response with a bad request error.
// Next, it retrieves the user details from the repository based on the provided email.
// If there is an error retrieving the user details, it returns a JSON response with an internal server error.
// If the user account is already active, it returns a JSON response with an internal server error indicating that the account is already verified.
// Otherwise, it creates a verification link token and sends an email to the user's email address.
// The function then updates the user's device information in the database.
// Finally, it returns a pointer to a models.RegistrationResponse object containing the user ID, email, and verification token.
func ResendVerificationLink(c *gin.Context) *models.RegistrationResponse {
	//* Get data for body
	reqBody := models.BodyRegisterRequest{}

	//* Check body valid
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	//* Get detail users
	resultDetailUser, err := repo.GetUserDetail(global.DB, reqBody.Email)

	log.Print(err)

	// * Check account exit into yet
	if err != nil {
		errorDetailUser := utils.HandleDBError(err)
		//* Error for database
		if errorDetailUser != "" {
			c.JSON(response.StatusInternalServerError, response.InternalServerError(errorDetailUser))
			return nil
		}
	}

	if resultDetailUser.IsActive {
		c.JSON(response.StatusBadRequest, response.BadRequestError("Account is verification"))
		return nil
	}

	resultVerificationLink := createTokenVerificationLink(c, models.UserIDEmail{
		ID:    resultDetailUser.ID,
		Email: reqBody.Email,
	})

	//* Send email
	data := models.EmailData{
		Title:    "Resend Verification User!",
		Body:     resultVerificationLink.Link,
		Template: `<h1>{{.Title}}</h1>Verification account: <a href="{{.Body}}">Click here to verify your account</a> </br> <img src="cid:logo" alt="Image" height="200" />`,
	}

	upsetDevice(c, resultDetailUser.ID, "")

	go pkg.SendGoEmail(reqBody.Email, data)

	return &models.RegistrationResponse{
		ID:    resultDetailUser.ID,
		Email: reqBody.Email,
		Token: resultVerificationLink.Token,
	}
}

// createKeyAndToken generates a random key pair, encodes the public key to PEM format,
// and creates access and refresh tokens using the provided user information and private key.
// It returns the access token, refresh token, and encoded public key.
func createKeyAndToken(resultUser models.UserIDEmail) (string, string, string) {
	privateKey, publicKey, err := helpers.RandomKeyPair()

	if err != nil {
		return "", "", ""
	}

	resultEncodePublicKey := helpers.EncodePublicKeyToPem(publicKey)

	accessToken, err := helpers.CreateToken(models.Payload{
		ID:    resultUser.ID,
		Email: resultUser.Email,
	}, privateKey, constants.ExpiresAccessToken)

	if err != nil {
		return "", "", ""
	}

	refetchToken, err := helpers.CreateToken(models.Payload{
		ID:    resultUser.ID,
		Email: resultUser.Email,
	}, privateKey, constants.ExpiresRefreshToken)

	if err != nil {
		return "", "", ""
	}

	return accessToken, refetchToken, resultEncodePublicKey
}

// setCookie sets a cookie in the response with the specified name, value, path, and maxAge.
// The domain is determined based on the environment and the request's host.
// The secure flag is set based on whether the environment is not the development environment.
// The httpOnly flag is set to true if the environment is not the development environment.
func setCookie(c *gin.Context, name string, value string, path string, maxAge int) {
	// Set up environment-related variables
	nodeEnv := os.Getenv("ENV")
	domain := global.Cfg.Server.Host
	secure := nodeEnv != constants.DevEnvironment
	httpOnly := false

	if nodeEnv != constants.DevEnvironment {
		hostWithPort := c.Request.Host
		parts := strings.Split(hostWithPort, ":")
		domain = parts[0]
		httpOnly = true
	}

	// Create a new cookie
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		Secure:   secure,
		HttpOnly: httpOnly,
	}

	// Set the cookie in the response
	http.SetCookie(c.Writer, cookie)
}

// upsetDevice updates or inserts a new device record in the database for the given user.
// It takes a gin.Context, user ID, and encoded public key as input parameters.
// It returns a pointer to the updated device information if successful, otherwise it returns nil.
func upsetDevice(c *gin.Context, id int, resultEncodePublicKey string) *models.Device {
	deviceID, _ := c.Get("device_id")

	ip := sql.NullString{String: c.ClientIP(), Valid: true}

	var publicKey string
	if resultEncodePublicKey != "" {
		publicKey = resultEncodePublicKey
	} else {
		publicKey = ""
	}

	resultInfoDevice, err := repo.UpsetDevice(global.DB, models.UpsetDeviceParams{
		UserID:     id,
		DeviceID:   deviceID.(string),
		DeviceType: c.Request.UserAgent(),
		Ip:         ip,
		PublicKey:  publicKey,
	})

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}
	return &resultInfoDevice
}

// createTokenVerificationLink generates a token verification link for the given user.
// It creates a random token for user verification, links the token with the user,
// and returns a TokenVerificationLink containing the token and the verification link.
// If any error occurs during token generation or database operations, it returns nil.
// The function takes a gin.Context and a user models.UserIDEmail as parameters.
func createTokenVerificationLink(c *gin.Context, user models.UserIDEmail) *models.TokenVerificationLink {
	//* Random Token for user verification
	token, err := helpers.GenerateToken()
	ExpiresAtToken := time.Now().Add(24 * time.Hour)
	ExpiresAtTokenUnix := ExpiresAtToken.Unix()

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	//* Link token with user
	linkVerification := fmt.Sprintf("%s/create/account/%s/%s/%s/%s", global.Cfg.Server.PortFrontend, user.Email, strconv.FormatInt(ExpiresAtTokenUnix, 10), strconv.Itoa(user.ID), token)

	verification := models.BodyVerificationRequest{
		UserId:        user.ID,
		VerifiedToken: token,
		ExpiresAt:     ExpiresAtToken,
	}

	_, err = repo.CreateVerification(global.DB, verification)

	if err != nil {
		//* Error for database
		errorCreateVerification := utils.HandleDBError(err)
		if errorCreateVerification != "" {
			c.JSON(response.StatusInternalServerError, response.InternalServerError(errorCreateVerification))
			return nil
		}
		c.JSON(response.StatusBadRequest, response.InternalServerError())
		return nil
	}

	return &models.TokenVerificationLink{
		Token: token,
		Link:  linkVerification,
	}

}

// checkUserIsActive checks if a user is active.
// If the user is not active, it returns nil.
// Otherwise, it returns a pointer to the isActive parameter.
func checkUserIsActive(isActive bool) *bool {
	if !isActive {
		return nil
	}
	return &isActive
}
