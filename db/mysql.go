package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@/test")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("数据库测试连接失败:", err)
	}
}
