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
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo/redis"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	pkg "github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/mail"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register handles the registration process for a user.
// It checks for user spam, validates the request body, checks if the user already exists,
// creates a new user if not, generates a verification link, sends an email for verification,
// and returns the registration response containing the user ID, email, and verification token.
// If any error occurs during the process, it returns an appropriate error response.
func Register(c *gin.Context) *models.RegistrationResponse {
	// * Check UserSpam
	resultSpam := redis.SpamUser(c, global.Cache, constants.SpamKey, constants.RequestThreshold)

	if resultSpam.IsSpam {
		ttl := fmt.Sprintf("You are blocked for %d seconds", resultSpam.ExpiredSpam)
		c.JSON(response.StatusBadRequest, response.BadRequestError(ttl))
		return nil
	}

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
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	ExpiresAtToken := time.Now().Add(24 * time.Hour)

	resultVerificationLink := createTokenVerificationLink(c, models.UserIDEmail{
		ID:    resultCreateUser.ID,
		Email: reqBody.Email,
	}, constants.StatusRegister, ExpiresAtToken)

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

// LoginIdentifier handles the login process for the user. It takes a gin.Context object as a parameter and returns a pointer to a models.LoginResponse struct.
// The function first checks if the user is marked as spam based on the request threshold. If the user is marked as spam, it returns an error response.
// Then, it checks the validity of the request body and binds it to the reqBody variable. If the request body is invalid, it returns an error response.
// Next, it identifies the type of identifier (email, phone, or username) and retrieves the user information from the database based on the identifier.
// If no user is found for the given identifier, it returns an error response.
// After that, it checks if the user account is active. If the account is blocked, it returns a forbidden error response.
// Then, it compares the provided password with the hashed password stored in the database. If the passwords do not match, it returns an error response.
// If the password is correct, it creates an access token, a refresh token, and encodes the public key.
// If any of the tokens or the encoded public key is empty, it returns an error response.
// It then updates the user's device information in the database and sets a cookie with the refresh token.
// Finally, it returns a LoginResponse object containing the user ID, device ID, email, and access token.
func LoginIdentifier(c *gin.Context) *models.LoginResponse {
	// * Check UserSpam
	resultSpam := redis.SpamUser(c, global.Cache, constants.SpamKeyLogin, constants.RequestThreshold)

	if resultSpam.IsSpam {
		ttl := fmt.Sprintf("You are blocked for %d seconds", resultSpam.ExpiredSpam)
		c.JSON(response.StatusBadRequest, response.BadRequestError(ttl))
		return nil
	}

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

	errPassword := helpers.ComparePassword(reqBody.Password, resultUser.PasswordHash.String)

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

// ResendVerificationLink is a function that handles the resend verification link request.
// It takes a gin.Context object as a parameter and returns a pointer to a models.RegistrationResponse object.
// The function first retrieves the request body data and checks its validity.
// If the request body is invalid, it returns a bad request error response.
// It then checks if the user is marked as spam based on the verification link requests made.
// If the user is marked as spam, it returns a bad request error response with a message indicating the blocking duration.
// Next, it retrieves the details of the user based on the provided email.
// If there is an error retrieving the user details, it returns an appropriate error response.
// It then counts the number of verification link requests made by the user.
// If the user has reached the maximum number of allowed verification link requests, it returns a bad request error response.
// If the user's account is already active, it returns a bad request error response indicating that the account is already verified.
// It then generates a verification link token for the user and sends an email containing the verification link.
// Finally, it returns a registration response object containing the user ID, email, and verification token.
func ResendVerificationLink(c *gin.Context) *models.RegistrationResponse {
	//* Get data for body
	reqBody := models.BodyRegisterRequest{}

	//* Check body valid
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}
	// * Check UserSpam
	resultSpam := redis.SpamUser(c, global.Cache, constants.SpamKeyLinkVerification, constants.RequestThresholdLinkVerification)

	if resultSpam.IsSpam {
		ttl := fmt.Sprintf("You are blocked for %d seconds", resultSpam.ExpiredSpam)
		c.JSON(response.StatusBadRequest, response.BadRequestError(ttl))
		return nil
	}

	//* Get detail users
	resultDetailUser, err := repo.GetUserDetail(global.DB, reqBody.Email)
	log.Print("OK", err, resultDetailUser)

	// * Check account exit into yet
	if err != nil {
		errorDetailUser := utils.HandleDBError(err)
		//* Error for database
		if errorDetailUser != "" {
			c.JSON(response.StatusInternalServerError, response.InternalServerError(errorDetailUser))
			return nil
		}
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	//* Count user had send verification
	count, err := repo.GetVerificationByUserId(global.DB, resultDetailUser.ID)

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	log.Print(count, err)
	numberSend := 5
	if count >= numberSend {
		times := fmt.Sprintf("You have sent verification %d times", numberSend)
		c.JSON(response.StatusBadRequest, response.BadRequestError(times))
		return nil
	}

	if resultDetailUser.IsActive {
		c.JSON(response.StatusBadRequest, response.BadRequestError("Account is verification"))
		return nil
	}

	resultVerificationLink := createTokenVerificationLink(c, models.UserIDEmail{
		ID:    resultDetailUser.ID,
		Email: reqBody.Email,
	}, constants.StatusResend, time.Now().Add(24*time.Hour))

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

// ForgetPassword handles the forget password functionality.
// It checks if the user is marked as spam, binds the request body,
// retrieves the user details from the database, and sends a forget password link via email.
// If successful, it returns the user ID, email, and forget password token.
// If any error occurs, it returns nil.
func ForgetPassword(c *gin.Context) *models.ForgetResponse {
	// * Check UserSpam
	resultSpam := redis.SpamUser(c, global.Cache, constants.SpamKeyForget, constants.RequestThresholdForget)

	if resultSpam.IsSpam {
		ttl := fmt.Sprintf("You are blocked for %d seconds", resultSpam.ExpiredSpam)
		c.JSON(response.StatusBadRequest, response.BadRequestError(ttl))
		return nil
	}

	reqBody := models.BodyForgetRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	resultDetailUser, err := repo.GetUserDetail(global.DB, reqBody.Email)

	if err != nil {
		errorDetailUser := utils.HandleDBError(err)
		//* Error for database
		if errorDetailUser != "" {
			c.JSON(response.StatusInternalServerError, response.InternalServerError(errorDetailUser))
			return nil
		}
	}

	if !resultDetailUser.IsActive {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	ExpiresAtToken := time.Now().Add(15 * time.Minute)
	resultForgetLink := createTokenVerificationLink(c, models.UserIDEmail{
		ID:    resultDetailUser.ID,
		Email: reqBody.Email,
	}, constants.StatusForget, ExpiresAtToken)

	//* Send email
	data := models.EmailData{
		Title:    "Forget Password!",
		Body:     resultForgetLink.Link,
		Template: `<h1>{{.Title}}</h1>Forget account: <a href="{{.Body}}">Click here to reset password your account</a> </br> <img src="cid:logo" alt="Image" height="200" />`,
	}

	go pkg.SendGoEmail(reqBody.Email, data)

	return &models.ForgetResponse{
		Id:    resultDetailUser.ID,
		Email: resultDetailUser.Email,
		Token: resultForgetLink.Token,
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
func createTokenVerificationLink(c *gin.Context, user models.UserIDEmail, status int, expiresToken time.Time) *models.TokenVerificationLink {
	//* Random Token for user verification
	token, err := helpers.GenerateToken()
	ExpiresAtTokenUnix := expiresToken.Unix()

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	//* Link token with user
	var linkVerification string
	if status == constants.StatusRegister || status == constants.StatusResend {
		linkVerification = fmt.Sprintf("%s/create/account/%s/%s/%s/%s", global.Cfg.Server.PortFrontend, user.Email, strconv.FormatInt(ExpiresAtTokenUnix, 10), strconv.Itoa(user.ID), token)
	} else {
		linkVerification = fmt.Sprintf("%s/reset/password/%s/%s/%s", global.Cfg.Server.PortFrontend, strconv.FormatInt(ExpiresAtTokenUnix, 10), strconv.Itoa(user.ID), token)
	}

	verification := models.BodyVerificationRequest{
		UserId:        user.ID,
		VerifiedToken: token,
		ExpiresAt:     expiresToken,
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
