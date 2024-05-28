package routers

import (
	"net/http"
	"os"

	constants "github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	nodeEnv := os.Getenv("ENV")

	if nodeEnv != constants.DevEnvironment {
		gin.SetMode(gin.ReleaseMode)
	}

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
