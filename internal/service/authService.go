package service

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) *string {
	a := 10
	b := "Ok"
	if a == 10 {
		c.JSON(response.StatusBadRequest, response.BadRequestError())
		return nil
	}

	return &b
}
