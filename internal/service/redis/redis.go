package redis

import (
	"Zchat/pkg/log"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

func CreateClient() *redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "number",
			DB:    0,
		})
	}

	return redisClient
}

func SetKeyEx(key string, value string, timeout time.Duration) error {
	err := CreateClient().Set(ctx, key, value, timeout).Err()
	if err != nil {
		log.GetZapLogger().Error(err.Error())
		return err
	}

	return err
}

func GetKey(key string) (string, error) {
	value, err := CreateClient().Get(ctx, key).Result()
	if err == redis.Nil {
		log.GetZapLogger().Error("key has expired")
		return "", err
	}

	return value, err
}