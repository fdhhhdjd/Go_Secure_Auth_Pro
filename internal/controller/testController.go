package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	author := "Nguyen Tien Tai123311144446333"
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + "----" + author,
	})
}
