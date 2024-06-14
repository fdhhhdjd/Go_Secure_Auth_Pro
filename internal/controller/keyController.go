package controller

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

// GetCsRfToken generates a CSRF token and sends it as a response.
// It takes a gin.Context object as a parameter.
// Returns an error if any.
func GetCsRfToken(c *gin.Context) error {
	token := csrf.GetToken(c)
	response.Ok(c, "Get Key Token", token)
	return nil
}
