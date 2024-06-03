package controller

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/service"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

func GetProfileUser(c *gin.Context) error {
	result := service.GetProfileUser(c)
	if result == nil {
		return nil
	}
	response.Ok(c, "Get Profile User", result)
	return nil
}
