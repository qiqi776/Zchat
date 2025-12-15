package redis

import (
	"Zchat/config"
	"Zchat/pkg/zlog"
	"context"
	"time"
	"strconv"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

func init() {
	conf := config.GetConfig()
	host := conf.RedisConfig.Host
	port := conf.RedisConfig.Port
	password := conf.RedisConfig.Password
	db := conf.Db
	addr := host + ":" + strconv.Itoa(port)

	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func SetKeyEx(key string, value string, timeout time.Duration) error {
	err := redisClient.Set(ctx, key, value, timeout).Err()
	if err != nil {
		zlog.Error(err.Error())
		return err
	}
	return nil
}

func GetKey(key string) (string, error) {
	value, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		zlog.Error("key_with_expire has expired")
		return "", err
	}
	return value, nil
}