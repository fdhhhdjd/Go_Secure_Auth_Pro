package global

import (
	"database/sql"
	"fmt"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/controller/initialization"
)

var (
	Cfg configs.Config
	DB  *sql.DB
)

func init() {
	//* CONFIG
	var err error
	Cfg, err = configs.LoadConfig("configs")
	if err != nil {
		panic(err)
	}

	//* DATABASE
	DB, err = initialization.Connect(Cfg)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		panic(err)
	}

}
