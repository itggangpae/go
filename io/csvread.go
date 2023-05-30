package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 파일 오픈
	file, _ := os.Open("./output.csv")

	// csv reader 생성
	rdr := csv.NewReader(bufio.NewReader(file))

	// csv 내용 모두 읽기
	rows, _ := rdr.ReadAll()

	// 행,열 읽기
	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s ", rows[i][j])
		}
		fmt.Println()
	}
}
