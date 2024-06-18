package global

import (
	"database/sql"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/controller/initialization"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	pkg "github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/setting"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

var (
	Cfg          models.Config
	DB           *sql.DB
	Cache        *redis.Client
	AdminSdk     *firebase.App
	MessageQueue *amqp.Connection
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

	//* FIREBASE
	AdminSdk, err = pkg.InitializeApp()
	if err != nil {
		fmt.Printf("Error connecting to firebase: %v\n", err)
		panic(err)
	}

	//* RABBITMQ
	MessageQueue, err = initialization.ConnectRabbitMQ(Cfg.RabbitMQ.URL)
	if err != nil {
		fmt.Printf("Error connecting to RabbitMq: %v\n", err)
		panic(err)
	}
}
