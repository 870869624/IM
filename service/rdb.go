package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"wechat/db/redis"
	db "wechat/db/sqlc"
)

func InsertMannager(user db.User) error {
	rdb := redis.ConnectRdb()
	string, err := rdb.Set(context.Background(), user.Account, "online", time.Duration(0)).Result()
	fmt.Println(string, "成功存入redis")
	if err != nil {
		return err
	}
	return nil
}

func GetUser(key string) (string, error) {
	rdb := redis.ConnectRdb()
	string, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	//这一步暂时有点多余
	// user, err := JsonUnmarshal(userMessage)
	// if err != nil {
	// 	return db.User{}, err
	// }
	return string, nil
}

// 从redis里面移除在线状态，需要补齐心跳检测或者使用延时检测移除
func DeleteUser(key string) (int64, error) {
	rdb := redis.ConnectRdb()
	result, err := rdb.Del(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

type RdbMessage struct {
	From_account_id string
	To_account_id   string
	Content         string
}

// 存储离线消息
func InsertMessage(from_account_id string, content string, to_account_id string) (string, error) {
	rdb := redis.ConnectRdb()
	data := &RdbMessage{
		From_account_id: from_account_id,
		To_account_id:   to_account_id,
		Content:         content,
	}
	dataString, err := JsonMarshal(*data)
	if err != nil {
		return "", err
	}

	err = rdb.LPush(context.Background(), to_account_id, dataString).Err()
	if err != nil {
		return "", err
	}
	return string("存储成功"), nil
}

// 获取离线消息
func GetMessage(account string) (*[]RdbMessage, error) {
	rdb := redis.ConnectRdb()
	data, err := rdb.LRange(context.Background(), account, 0, -1).Result()
	outlineMessage := make([]RdbMessage, len(data))
	if err != nil {
		return &[]RdbMessage{}, err
	}
	if len(data) == 0 {
		return &[]RdbMessage{}, fmt.Errorf("123456789")
	}
	for i := 0; i < len(data); i++ {
		err = json.Unmarshal([]byte(data[i]), &outlineMessage[i])
		if err != nil {
			return &[]RdbMessage{}, err

		}
	}
	return &outlineMessage, nil
}
