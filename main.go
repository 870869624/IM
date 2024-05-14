package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"wechat/api"
	db "wechat/db/sqlc"
	"wechat/tcp"

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

	store := db.NewStore(connection)    //初始化,将数据库连接存入存储器
	server, err := api.NewServer(store) //利用存储器新建网络服务
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

	server1, err := tcp.NewServer(store)
	if err != nil {
		log.Fatal("tcp服务未能建立", err)
	}
	tcp.AcceptListen(listener, server1)

}
