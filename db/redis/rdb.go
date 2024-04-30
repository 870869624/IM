package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func ConnectRdb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
	return rdb
}
