package controllers

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// Pong is a handler function that returns a pong response.
// It takes a gin.Context object as a parameter and sends a response with the author's name.
func Pong(c *gin.Context) {
	author := "pong" + "----" + "Nguyen Tien Tai"
	response.Ok(c, "Test", author)
}
