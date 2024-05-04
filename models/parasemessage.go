package models

import (
	"context"
	"time"
	"wechat/db/redis"
	db "wechat/db/sqlc"
)

func UserMannager(user db.User) string {
	rdb := redis.ConnectRdb()
	_, err := rdb.Set(context.Background(), user.Account, user, time.Duration(0)).Result()
	if err != nil {
		return err.Error()
	}
	return ""
}

func GetUser(acount string) (db.User, error) {
	rdb := redis.ConnectRdb()
	userMessage, err := rdb.Get(context.Background(), acount).Result()
	if err != nil {
		return db.User{}, nil
	}
	user, err := JsonUnmarshal(userMessage)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
