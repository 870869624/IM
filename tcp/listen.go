package tcp

import (
	"fmt"
	"net"
)

func AcceptListen(listener net.Listener, server *Server) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("启动出错")
			return
		}
		//检测离线消息里面是否有离线消息

		go Process(conn, server)
	}
}
