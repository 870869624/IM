package service

import (
	"encoding/json"
	db "wechat/db/sqlc"
)

// 传入User信息返回解码出来的结构体(暂时搁置，没用上)
func JsonUnmarshal(user string) (*db.User, error) {
	var userMessage db.User
	err := json.Unmarshal([]byte(user), &userMessage)
	if err != nil {
		return &db.User{}, err
	}
	return &userMessage, nil
}

func JsonMarshal(rdbMessage RdbMessage) (string, error) {
	byte, err := json.Marshal(rdbMessage)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}
