package db

import (
	"context"
	"webback/config"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	RDB *redis.Client
	ctx = context.Background()
)

func InitRedis() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Address,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Database,
	})

	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("redis连接失败, err: %v", err)
		return
	}

	// LoadWebVisitorFromRedis() // temp

	return
}
