package main

import (
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/routers"
)

// @title 	Server Auth
// @version	1.0
// @description This is server auth API in Go using Gin framework
// @host 103.82.195.138:8000
// @BasePath /v1
func main() {
	r := routers.NewRouter()
	r.Run(":" + global.Cfg.Server.Port)
}
