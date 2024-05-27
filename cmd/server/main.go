package main

import "github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/routers"

func main() {
	r := routers.NewRouter()
	r.Run(":8002")
}
