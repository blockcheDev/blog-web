package db

import (
	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	// ctx = context.Background()
)

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //没有密码
		DB:       0,  // use default DB
	})
}
