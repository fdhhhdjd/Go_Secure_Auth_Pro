package controller

import (
	"net/http"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": service.Register(),
	})
}
