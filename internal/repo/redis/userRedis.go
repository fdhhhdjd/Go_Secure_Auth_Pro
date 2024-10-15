package redis

import (
	"context"
	"log"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// SpamUser checks if a user is spamming based on the request threshold and Cuckoo filter.
func SpamUser(ctx *gin.Context, rdb *redis.Client, key string, requestThreshold int64) *models.SpamUserResponse {
	numberRequest, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return nil
	}

	var ttl time.Duration
	if numberRequest == requestThreshold+1 {
		rdb.Expire(ctx, key, constants.InitialBlock)
		ttl = constants.InitialBlock
	} else if numberRequest > requestThreshold+2 {
		rdb.Expire(ctx, key, constants.ExtendedBlock)
		ttl = constants.ExtendedBlock
	} else {
		ttl, err = rdb.TTL(ctx, key).Result()
		if err != nil {
			return nil
		}
	}

	if numberRequest > requestThreshold {
		return &models.SpamUserResponse{
			ExpiredSpam: int(ttl.Seconds()),
			IsSpam:      true,
		}
	}

	rdb.Expire(ctx, key, constants.ExpireDuration)

	return &models.SpamUserResponse{
		ExpiredSpam: 0,
		IsSpam:      false,
	}
}

// AddUserToCuckooFilter adds a user to the Cuckoo filter in Redis and sets an expiration time.
func AddUserToCuckooFilter(ctx context.Context, rdb *redis.Client, key string, expiration time.Duration) error {
	cuckooKey := "cuckoo:" + key
	_, err := rdb.Do(ctx, "CF.ADD", cuckooKey, key).Result()
	if err != nil {
		return err
	}

	// Set expiration time
	_, err = rdb.Expire(ctx, cuckooKey, expiration).Result()
	return err
}

// DeleteKeyUser deletes the entire key from Redis.
func DeleteKeyUser(ctx context.Context, rdb *redis.Client, key string) error {
	log.Println("Deleting key from Redis: ", key)
	_, err := rdb.Del(ctx, key).Result()
	return err
}

func GetUserToCuckooFilter(ctx context.Context, rdb *redis.Client, key string) (bool, error) {
	cuckooKey := "cuckoo:" + key
	exists, err := rdb.Do(ctx, "CF.EXISTS", cuckooKey, key).Result()
	if err != nil {
		return false, err

	}

	log.Println("Exists: ", exists)
	return exists != 0, nil
}
