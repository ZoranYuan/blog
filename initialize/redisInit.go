package initialize

import (
	"ZoranYuan_blog/global"
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (*redis.Client, error) {
	redisConfig := global.Config.Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}
