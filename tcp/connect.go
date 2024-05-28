package tcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"
	db "wechat/db/sqlc"
	"wechat/mq"
	"wechat/service"
	"wechat/util"
)

// 连接成功以后把用户网络状态信息存入redis中
// 用户在线状态信息
//
//	type onlineMap struct{
//		user db.User
//	}

type Messages struct {
	Message int `json:"message"`
	Apply   int `json:"apply"`
	System  int `json:"system"`
}
type MessageType struct {
	From_account_account string   `json:"from_account_account"`
	To_account_account   string   `json:"to_account_account"`
	Group_account        string   `json:"group_account"`
	Token                string   `json:"token"`
	Type                 Messages //1.个人消息 2.群消息 3.系统消息 申请消息：群申请，个人申请
	Content              string   `json:"content"`
}

// 用户在线状态信息，key为account， val为连接
var OnlineMap = make(map[string]net.Conn)

// 获取输入信息，对比tokenString然后获取数据库用户信息，存入onlineMapz中,在这里处理在线离线，在connect中处理连接用户
func Process(conn net.Conn, server *Server) {
	var message []string
	defer conn.Close()
	var messageReq MessageType
	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err, time.Now())
			return
		}

		// 存入消息队列----------
		err = mq.PushMessage(data, "message")
		if err != nil {
			fmt.Println(err)
			return
		}

		msgs := mq.ConsumMessage("message")
		m := <-*msgs
		fmt.Println(msgs, "=====", m)
		for {
			fmt.Println(string(m.Body), "----------------")
			message = append(message, string(m.Body))
			fmt.Println(message, "+++++++++++++++")

			//json格式解码
			err = json.Unmarshal(data[:n], &message)

			// fmt.Printf("%+v", messageReq)   这里是用来检测收到消息格式有误问题的
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
			if claims.Account != messageReq.From_account_account { //token错误也会被中断，获取到用户的信息
				conn.Write([]byte("账户信息不正确"))
				conn.Close()
				return
			}
			user, err := server.store.GetUser(context.Background(), claims.Account) //获取到了用户信息和消息信息
			if err != nil {
				conn.Write([]byte(string(err.Error())))
				conn.Close()
				return
			}

			//检测是否有属于该用户的离线消息
			outlineMessage, err := service.GetMessage(claims.Account)
			if err == nil {
				for _, v := range *outlineMessage {
					_, err := conn.Write([]byte("  发送者:" + v.From_account_id + "  信息是:" + v.Content + "\n"))
					if err != nil {
						fmt.Println("离线消息获取失败")
					}
				}
			} else {
				fmt.Println("没有离线私人消息")
			}
			m := CheckOutlineGroupMessage(server, claims.Account) //群组离线消息
			if m != nil {
				for _, v := range m {
					_, err := conn.Write([]byte("  发送者:" + v.FromUserAccount + "  群组:" + v.GroupAccount + "  文本是:" + v.Content + "\n"))
					if err != nil {
						fmt.Println("离线消息获取失败")
					}
				}
			}

			go heartBeat() //简单的心跳检测
			//检查该登陆用户在线状态
			if !CheckOnline(user, conn) {
				OnlineMap[claims.Account] = conn //新增在线用户放入map中
			}

			fmt.Println("发送给:"+messageReq.To_account_account, "文本信息是:"+messageReq.Content, "----------")

			//获取消息结构体中的发送目标
			checkMessage(messageReq, OnlineMap, server)
		}

	}
}

// 心跳检测，因为要写客户端，暂时先写成延时检测,（并没有考虑到多地登陆，如果多地登陆的话应该需要加入其它参数）
func heartBeat() {
	for {
		for account, conn := range OnlineMap {
			_, err := conn.Write([]byte(""))
			if err != nil {
				fmt.Println(account + "已经断开连接")
				delete(OnlineMap, account)
				service.DeleteUser(account)
			}
		}
		time.Sleep(time.Second)
	}
	// defer models.DeleteUser(account) //移除用户在线状态
	// defer conn.Close()
	// timer := time.NewTimer(5 * time.Second)
	// defer timer.Stop()
	// fmt.Println("---", conn.RemoteAddr().String())
	// for {
	// 	select {
	// 	case <-timer.C:
	// 		fmt.Println("客户端已经超时")
	// 		return
	// 	default:
	// 		conn.SetReadDeadline(time.Now().Add(20 * time.Second))
	// 		data := make([]byte, 1024)

	// 		n, err := conn.Read(data)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		if n > 0 {
	// 			timer.Reset(5 * time.Second)
	// 			fmt.Println(string(data[:n]))
	// 		}
	// 	}
	// }
}

// 处理连接用户， 发消息给用户，如果不在线就放入redis保存七天并且放入到mysql中（历史记录不需要存储是否为离线消息）
func connectUser(messageReq MessageType, server *Server) {
	connTo, ok := OnlineMap[messageReq.To_account_account]
	if ok {
		fmt.Println("检测到有对象用户")
		arg := db.CreateMessageParams{
			Content:         messageReq.Content,
			FromUserAccount: messageReq.From_account_account,
			ToUserAccount:   messageReq.To_account_account,
			GroupAccount:    messageReq.Group_account,
			MType:           1,
			Networkstatus:   1,
		}
		_, err := server.store.CreateMessage(context.Background(), arg)
		if err != nil {
			fmt.Println(err, "保存信息失败")
			return
		}

		connTo.Write([]byte(messageReq.Content + "\n"))
		return
	}

	fmt.Println("接收对象不在线,存入redis中")
	//如果在map里没有说明不在线存入离线消息，并且存入数据库
	sign, err := service.InsertMessage(messageReq.From_account_account, messageReq.Content, messageReq.To_account_account)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sign)
	//放入mysql中
	arg := db.CreateMessageParams{
		Content:         messageReq.Content,
		FromUserAccount: messageReq.From_account_account,
		ToUserAccount:   messageReq.To_account_account,
		GroupAccount:    messageReq.Group_account,
		MType:           1,
		Networkstatus:   0,
	}
	_, err = server.store.CreateMessage(context.Background(), arg)
	if err != nil {
		fmt.Println(err, "保存信息失败")
		return
	}

}

// 处理链接群组处理群消息
func connectGroup(messageReq MessageType, server *Server) {
	useraccounts, err := server.store.GetGMAccount(context.Background(), messageReq.Group_account) //获取对应群组的所有成员
	if err != nil {
		fmt.Println(useraccounts, "2222", err)
		return
	}
	if len(useraccounts) == 0 { //没有该群组
		fmt.Println("没有该群组")
	}

	//查看该用户是否在该群组中是否有发送信息的这个用户
	findUser := db.CheckUserParams{
		GroupAccount: messageReq.Group_account,
		UserAccount:  messageReq.From_account_account,
	}
	_, err = server.store.CheckUser(context.Background(), findUser)
	if err != nil {
		fmt.Println("该群组没有该用户")
		return
	}

	//判断在线map中是否有群组用户
	for _, v := range useraccounts {
		connTo, ok := OnlineMap[v]
		if ok {
			connTo.Write([]byte(messageReq.Content + "\n"))
		} else {
			fmt.Println(v + "不在线")
		}
	}

	arg := db.CreateMessageParams{
		Content:         messageReq.Content,
		FromUserAccount: messageReq.From_account_account,
		ToUserAccount:   messageReq.To_account_account,
		GroupAccount:    messageReq.Group_account,
		MType:           2,
		Networkstatus:   0,
	}
	_, err = server.store.CreateMessage(context.Background(), arg)
	if err != nil {
		fmt.Println(err, "保存信息失败")
		return
	}
}

func checkMessage(messageReq MessageType, OnlineMap map[string]net.Conn, server *Server) {
	fmt.Println("进入消息检测")
	if messageReq.To_account_account != "" {
		fmt.Println(OnlineMap, "进入对用户发消息检测")
		connectUser(messageReq, server)
		return
	}
	if messageReq.Group_account != "" {
		fmt.Println(OnlineMap, "进入对群组发消息检测")
		connectGroup(messageReq, server)
		return
	}
	fmt.Println("发送对象为空, 此条是登陆消息")
}
