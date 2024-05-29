package controller

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	author := "pong" + "----" + "Nguyen Tien Tai"
	response.Ok(c, "Test", author)
}
