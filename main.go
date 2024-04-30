package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"wechat/api"
	db "wechat/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://postgres:123456@localhost:5432/wechat?sslmode=disable"
)

func main() {
	//gin服务
	connection, err := sql.Open(DBDriver, DBSource)
	fmt.Println(DBDriver, DBSource)
	if err != nil {
		log.Fatalln("无法连接", err)
	}

	store := db.NewStore(connection) //初始化
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("不能新建服务", err)
	}
	go server.Strat(":3000")

	//tcp服务
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("启动出错")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("启动出错")
			return
		}

		go api.Process(conn, server)
	}

}