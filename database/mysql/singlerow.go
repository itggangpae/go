package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:wnddkd@tcp(127.0.0.1:3306)/adam")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
