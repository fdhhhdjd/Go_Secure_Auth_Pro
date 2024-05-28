package controller

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/service"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) error {
	result := service.Register(c)
	if result == nil {
		return nil
	}
	response.Created(c, "Register", result)
	return nil
}
