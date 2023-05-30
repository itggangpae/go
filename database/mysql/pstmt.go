package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:wnddkd@tcp(127.0.0.1:3306)/adam")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepared Statement 생성
	stmt, err := db.Prepare("UPDATE test1 SET name=? WHERE id=?")
	checkError(err)
	defer stmt.Close()

	// Prepared Statement 실행
	_, err = stmt.Exec("one", 1) //Placeholder 파라미터 순서대로 전달
	checkError(err)
	_, err = stmt.Exec("two", 2)
	checkError(err)
	_, err = stmt.Exec("three", 3)
	checkError(err)
}
