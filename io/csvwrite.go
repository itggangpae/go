package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

func main() {
	// 파일 생성
	file, err := os.Create("./output.csv")
	if err != nil {
		panic(err)
	}

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	wr.Write([]string{"adam", "53"})
	wr.Write([]string{"jessica", "51"})
	wr.Flush()
}
