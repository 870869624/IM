package models

import (
	"encoding/json"
	db "wechat/db/sqlc"
)

// 传入User信息返回解码出来的结构体
func JsonUnmarshal(user string) (db.User, error) {
	var userMessage db.User
	err := json.Unmarshal([]byte(user), &userMessage)
	if err != nil {
		return db.User{}, err
	}
	return userMessage, nil
}
