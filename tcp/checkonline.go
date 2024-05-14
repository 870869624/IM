package tcp

import (
	"context"
	"fmt"
	"net"
	db "wechat/db/sqlc"
	"wechat/service"
)

func CheckOnline(user db.User, conn net.Conn) bool {
	str, err := service.GetUser(user.Account)
	if err != nil {
		fmt.Println("redis没有在线状态记录,存入redis") //redis里没有所以进来了
		// 在线状态存储进redis
		service.InsertMannager(user) // 存入redis中
		return false                 //需要在map里保存
	}
	fmt.Println("当前该用户状态:", str)
	return true
}

// 写在了db里面
func CheckOutlineGroupMessage(server *Server, account string) []db.GetMessageToGroupRow {
	var m []db.GetMessageToGroupRow
	//先去查找该用户所有加入的群组, 然后获取群组account之后去message里面，获取status为离线且account对应的群组
	groupAccounts, err := server.store.GetUserGroup(context.Background(), account)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(groupAccounts) == 0 {
		fmt.Println("用户没有加入任何群组")
		return nil
	}
	for _, v := range groupAccounts {
		message, err := server.store.GetMessageToGroup(context.Background(), v)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if len(message) == 0 {
			fmt.Println("没有群组离线消息")
			return nil
		}
		m = append(m, message...)
	}
	return m
}
