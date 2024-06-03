package service

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/repo"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

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
