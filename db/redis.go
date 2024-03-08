package db

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	RDB *redis.Client
	ctx = context.Background()
)

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //没有密码
		DB:       0,  // use default DB
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("redis连接失败, err: %v", err)
	}
}
