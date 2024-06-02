package redis

import (
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

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
