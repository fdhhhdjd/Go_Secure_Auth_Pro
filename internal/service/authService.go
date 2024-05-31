package service

import (
	"fmt"
	"log"
	"strconv"
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

	//* Random Token for user verification
	token, err := utils.GenerateToken()
	ExpiresAtToken := time.Now().Add(24 * time.Hour)
	ExpiresAtTokenUnix := ExpiresAtToken.Unix()

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	//* Link token with user
	linkVerification := fmt.Sprintf("%s/create/account/%s/%s/%s/%s", global.Cfg.Server.PortFrontend, reqBody.Email, strconv.FormatInt(ExpiresAtTokenUnix, 10), strconv.Itoa(resultCreateUser.ID), token)

	verification := models.BodyVerificationRequest{
		UserId:        resultCreateUser.ID,
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

	//* Send email
	data := models.EmailData{
		Title:    "Register User!",
		Body:     linkVerification,
		Template: `<h1>{{.Title}}</h1>Verification account: <a href="{{.Body}}">Click here to verify your account</a> </br> <img src="cid:logo" alt="Image" height="200" />`,
	}

	go pkg.SendGoEmail(reqBody.Email, data)

	return &models.RegistrationResponse{
		ID:    resultCreateUser.ID,
		Email: reqBody.Email,
		Token: token,
	}
}

// VerificationAccount is a function that handles the verification of a user's account.
// It takes a Gin context as input and returns a VerificationResponse pointer.
// The function first binds the query parameters from the context to a QueryLoginRequest struct.
// If there is an error in binding the query parameters, it returns a BadRequestError response.
// It then retrieves the verification details from the database using the GetVerification function.
// If there is an error in retrieving the verification details, it returns a BadRequestError response.
// Next, it checks if the retrieved verification details match the query parameters and if the verification is active.
// If any of the conditions fail, it returns a BadRequestError response.
// It also checks if the verification token has expired. If it has, it returns an UnauthorizedError response.
// If all the checks pass, it generates a random password, hashes it, and updates the user's password in the database.
// It also inserts the old password into the password history table.
// Then, it updates the verification status of the user in the database.
// After that, it sends an email to the user with the new password.
// Finally, it returns a VerificationResponse with the verification details.
func VerificationAccount(c *gin.Context) *models.VerificationResponse {
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
	})
	log.Print(errUpdatePassword)
	if errUpdatePassword != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	//* Send email
	data := models.EmailData{
		Title:    "Verification Account Success!",
		Body:     randomPassword,
		Template: `<h1>{{.Title}}</h1> <p style="font-size: large;">Thank You, You have verification account success 😊. <br/> Password New: <b>{{.Body}}</b></p>`,
	}

	go pkg.SendGoEmail(resultUpdateUser.Email, data)

	return &models.VerificationResponse{
		ID:     GetVerification.ID,
		UserId: GetVerification.UserID,
		Token:  GetVerification.VerifiedToken,
	}
}

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

	errPassword := helpers.ComparePassword(reqBody.Password, resultUser.PasswordHash)

	if errPassword != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	privateKey, publicKey, err := helpers.RandomKeyPair()

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	resultEncodePublicKey := helpers.EncodePublicKeyToPem(publicKey)

	log.Print(resultEncodePublicKey)

	accessToken, err := helpers.CreateToken(models.Payload{
		ID:    resultUser.ID,
		Email: resultUser.Email,
	}, privateKey, 15*time.Minute)

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	refetchToken, err := helpers.CreateToken(models.Payload{
		ID:    resultUser.ID,
		Email: resultUser.Email,
	}, privateKey, 30*24*time.Hour)

	if err != nil {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	log.Print(refetchToken)

	return &models.LoginResponse{
		ID:          resultUser.ID,
		Email:       resultUser.Email,
		AccessToken: accessToken,
	}
}
