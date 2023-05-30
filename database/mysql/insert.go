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

	// INSERT 문 실행
	result, err := db.Exec("INSERT INTO test1 VALUES (?, ?)", 101, "yena")
	if err != nil {
		log.Fatal(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
}
