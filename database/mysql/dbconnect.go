package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:wnddkd@tcp(127.0.0.1:3306)/adam")
	if err != nil {
		log.Fatal(err)
	}
	if err == nil {
		log.Println("연결 성공")
	}
	defer db.Close()
}
