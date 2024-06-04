package initialization

import (
	"context"
	"fmt"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs"
	"github.com/redis/go-redis/v9"
)

// ConnectRedis establishes a connection to Redis using the provided configuration.
// It returns a Redis client and an error if the connection fails.
func ConnectRedis(cfg configs.Config) (*redis.Client, error) {
	var rdb *redis.Client
	var pong string
	var err error

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Cache.Host, cfg.Cache.Port),
			Password: cfg.Cache.Password,
			DB:       0,
		})

		pong, err = rdb.Ping(context.Background()).Result()
		if err != nil {
			fmt.Printf("Error connecting to Redis: %v\n", err)
			fmt.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		} else {
			fmt.Println("CONNECTED TO REDIS:", pong, "🥩")
			return rdb, nil
		}
	}

	return nil, fmt.Errorf("could not connect to Redis after %d attempts", maxRetries)
}
