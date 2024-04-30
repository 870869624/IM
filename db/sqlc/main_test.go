package db

import (
	"database/sql"
	"fmt"

	"testing"
)

var testQuires *Queries

const (
	dbDriver = "postgres "
	dbSource = "postgresql://postgres:123456@localhost:5432/wechat?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Println("连接超时")
		return
	}
	testQuires = New(conn)

	m.Run()
}
