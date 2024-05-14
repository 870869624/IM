package service

// import (
// 	"context"
// 	"fmt"
// 	db "wechat/db/sqlc"
// )

// type Messages struct {
// 	Message string `json:"message"`
// 	Apply   string `json:"apply"`
// 	System  string `json:"system"`
// }
// type MessageType struct {
// 	From_account_id string   `json:"from_account_id"`
// 	To_account_id   string   `json:"to_account_id"`
// 	Group_id        string   `json:"group_id"`
// 	Token           string   `json:"token"`
// 	Type            Messages //1.个人消息 2.群消息 3.系统消息 申请消息：群申请，个人申请
// 	Content         string   `json:"content"`
// }

// func DBCreateMessage(message MessageType) {
// 	arg := db.CreateMessageParams{
// 		Content:    messageReq.Content,
// 		FromUserID: messageReq.From_account_id,
// 		ToUserID:   messageReq.To_account_id,
// 		GroupID:    messageReq.Group_id,
// 		MType:      1,
// 	}
// 	_, err = server.store.CreateMessage(context.Background(), arg)
// 	if err != nil {
// 		fmt.Println(err, "保存信息失败")
// 		return
// 	}
// }
