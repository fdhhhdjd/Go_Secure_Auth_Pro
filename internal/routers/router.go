package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Pong)
	return r
}

func Pong(c *gin.Context) {
	author := "Nguyen Tien Tai"
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + "----" + author,
	})
}
