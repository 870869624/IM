package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"
	"wechat/models"
	"wechat/util"
)

// 连接成功以后把用户网络状态信息存入redis中
// 用户在线状态信息
//
//	type onlineMap struct{
//		user db.User
//	}

type Messages struct {
	Message string `json:"message"`
	Apply   string `json:"apply"`
	System  string `json:"system"`
}
type MessageType struct {
	From_account_id string   `json:"from_account_id"`
	To_account_id   string   `json:"to_account_id"`
	Group_id        string   `json:"group_id"` 
	Token           string   `json:"token"`
	Type            Messages //1.个人消息 2.群消息 3.系统消息 申请消息：群申请，个人申请
	Content         string   `json:"content"`
}

// 用户在线状态信息，key为account， val为连接
var OnlineMap map[string]net.Conn

// 获取输入信息，对比tokenString然后获取数据库用户信息，存入onlineMapz中,在这里处理在线离线，在connect中处理连接用户
func Process(conn net.Conn, server *Server) {
	defer conn.Close()

	var messageReq MessageType
	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err, time.Now())
			return
		}
		err = json.Unmarshal(data[:n], &messageReq)

		fmt.Printf("%+v", messageReq)
		if err != nil {
			conn.Write([]byte(string(err.Error())))
			conn.Close()
			return
		}

		claims, err := util.ParaseToken(messageReq.Token) //如果token过期请求就会被中断
		if err != nil {
			conn.Write([]byte(string(err.Error())))
			conn.Close()
			return
		}
		//token错误也会被中断，获取到用户的信息
		user, err := server.store.GetUser(context.Background(), claims.Account) //获取到了用户信息和消息信息
		if err != nil {
			conn.Write([]byte(string(err.Error())))
			conn.Close()
			return
		}

		//在线状态存储进redis
		models.UserMannager(user)
		OnlineMap[claims.Account] = conn

		//获取消息结构体中的发送目标
		checkMessage(messageReq, conn)
	}
}

// 心跳检测，因为要写客户端，暂时先写成延时检测
func heartBeat(conn net.Conn) {
	defer conn.Close()
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	fmt.Println("---", conn.RemoteAddr().String())
	for {
		select {
		case <-timer.C:
			fmt.Println("客户端已经超时")
			return
		default:
			conn.SetReadDeadline(time.Now().Add(20 * time.Second))
			data := make([]byte, 1024)

			n, err := conn.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			if n > 0 {
				timer.Reset(5 * time.Second)
				fmt.Println(string(data[:n]))
			}
		}
	}
}

// 处理连接用户
func connctUser(conn net.Conn) {

}

// 处理链接群组处理群消息
func connectGroup() {

}

func checkMessage(messageReq MessageType, conn net.Conn) {
	if messageReq.To_account_id != "" {
		for account, connTo := range OnlineMap {
			if account == messageReq.To_account_id {
				connTo.Write([]byte(messageReq.Content))
			}
		}
	}
}
