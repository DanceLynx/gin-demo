package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hello/config"
)

var Redis *redis.Client

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
		DB:   config.Redis.DB, // use default DB
	})
	ctx := context.Background()
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	InitLogger.Info("redis", "connnect to redis successful")
}
