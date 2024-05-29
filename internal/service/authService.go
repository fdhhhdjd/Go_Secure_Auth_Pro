package service

import (
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
// Next, it retrieves the user details from the database based on the provided email.
// If there is an error retrieving the details, it returns an InternalServerError response.
// If the user already exists, it returns a BadRequestError response.
// If the user does not exist, it creates a new user in the database.
// If there is an error creating the user, it returns an InternalServerError response.
// Finally, it returns a RegistrationResponse object containing the created user's ID and email.
// Register handles the registration process for a user.
// It receives a Gin context object and returns a RegistrationResponse pointer.
// The function first retrieves the request body data and checks its validity.
// If the body is invalid, it returns a BadRequestError response.
// Next, it retrieves the user details from the database based on the provided email.
// If there is an error retrieving the details, it returns an InternalServerError response.
// If the user already exists, it returns a BadRequestError response.
// If the user does not exist, it creates a new user in the database.
// If there is an error creating the user, it returns an InternalServerError response.
// Finally, it returns a RegistrationResponse object containing the created user's ID and email.
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

	//*  Send email
	data := models.EmailData{
		Title:    "Register User!",
		Body:     "This is some HTML content",
		Template: `<h1>{{.Title}}</h1><p>{{.Body}}</p><img src="cid:logo" alt="Image" height="200" />`,
	}

	pkg.SendGoEmail(reqBody.Email, data)

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

	return &models.RegistrationResponse{
		ID:    resultCreateUser.ID,
		Email: reqBody.Email,
	}
}
