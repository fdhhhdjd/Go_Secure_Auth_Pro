package global

import (
	"database/sql"
	"fmt"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/controller/initialization"
	"github.com/redis/go-redis/v9"
)

var (
	Cfg   configs.Config
	DB    *sql.DB
	Cache *redis.Client
)

func init() {
	//* CONFIG
	var err error
	Cfg, err = configs.LoadConfig("configs/yaml")
	if err != nil {
		panic(err)
	}

	//* DATABASE
	DB, err = initialization.ConnectPG(Cfg)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		panic(err)
	}

	//* CACHE
	Cache, err = initialization.ConnectRedis(Cfg)
	if err != nil {
		fmt.Printf("Error connecting to Redis: %v\n", err)
		panic(err)
	}
}
