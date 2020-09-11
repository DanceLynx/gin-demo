package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hello/config"
	"fmt"
	"runtime"
)

var Redis *redis.Client

func ConnectRedis() {
	redis.SetLogger(&redisLogger{})
	Redis = redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
		DB:   config.Redis.DB, // use default DB
	})
	ctx := context.Background()
	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	
	InitLogger.Info("connnect to redis successful")
}


type redisLogger struct{
}

func (this *redisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(3)
	message := fmt.Sprintf(format, v...)
	m := map[string]interface{} {
		"file":fmt.Sprintf("%s:%d", file, line),
		"message":message,
	}
	Logger.Warn(ctx,"redis",m)
}