package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() *sql.DB {
	//db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.10:3306)/gotest?charset=utf8")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/work?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

/**
  json格式文件定义
**/
type ReturnJson struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}
