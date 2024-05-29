package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/utils"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	pkg "github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/mail"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
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
	linkVerification := fmt.Sprintf("%s/create/account/%s/%s/%s", global.Cfg.Server.PortFrontend, strconv.FormatInt(ExpiresAtTokenUnix, 10), strconv.Itoa(resultCreateUser.ID), token)

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
