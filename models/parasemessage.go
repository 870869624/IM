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

func GetUser() (db.User, error) {

}
