package models

import (
	"context"
	"time"
	"wechat/db"
	"wechat/models"
)

func UserMannager(user models.User) error {
	rdb := db.ConnectRdb()
	err := rdb.Set(context.Background(), user.Account, user, time.Duration(0))
	if err != nil {
		return err.Err()
	}
	return nil
}
