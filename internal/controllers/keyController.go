package controllers

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// GetCsRfToken generates a CSRF token and sends it as a response.
// It takes a gin.Context object as a parameter.
// Returns an error if any.
// @Summary Get CSRF Token
// @Description Generates a CSRF token and sends it as a response
// @Tags CSRF
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param X-Device-Id header string true "Device ID"
// @Success 200 {string} string "CSRF token"
// @Router /key/csrf-token [get]
func GetCsRfToken(c *gin.Context) error {
	token := csrf.GetToken(c)
	response.Ok(c, "CSRF Token", token)
	return nil
}
