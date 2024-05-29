package initialization

import (
	"context"
	"fmt"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs"
	"github.com/redis/go-redis/v9"
)

// ConnectRedis establishes a connection to Redis using the provided configuration.
// It returns a Redis client and an error if the connection fails.
func ConnectRedis(cfg configs.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Cache.Host, cfg.Cache.Port),
		Password: cfg.Cache.Password,
		DB:       0,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	} else {
		fmt.Println("CONNECTED TO REDIS:", pong, "🥩")
	}
	return rdb, nil
}
