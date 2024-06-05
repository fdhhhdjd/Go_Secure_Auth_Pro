package service

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

func LoginSocial(c *gin.Context) *models.LoginResponse {
	reqBody := models.BodyLoginSocialRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.BadRequestError(c)
		return nil
	}

	infoUserSocial := helpers.GetUserIDToken(c, reqBody.Uid)

	if infoUserSocial == nil {
		response.BadRequestError(c)
		return nil
	}

	return &models.LoginResponse{}
}
