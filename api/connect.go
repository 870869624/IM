package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"
	"wechat/util"
)

type Messages struct {
	Message string `json:"message"`
	Apply   string `json:"apply"`
}
type MessageType struct {
	From_account_id string   `json:"from_account_id"`
	To_account_id   string   `json:"to_account_id"`
	Group_id        string   `json:"group_id"`
	Token           string   `json:"token"`
	Type            Messages //1.个人消息 2.群消息 3.系统消息 申请消息：群申请，个人申请
}

// 获取输入信息，对比tokenString然后获取数据库用户信息，存入onlineMapz中
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
		//token错误也会被中断
		user, err := server.store.GetUser(context.Background(), claims.Account)
		if err != nil {
			conn.Write([]byte(string(err.Error())))
			conn.Close()
			return
		}
		fmt.Println(claims, user)
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
