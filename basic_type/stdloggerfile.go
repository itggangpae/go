package main

import (
	"log"
	"os"
)

func main() {
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	// 표준로거를 파일로그로 변경
	log.SetOutput(fpLog)

	run()
	log.Println("프로그램 종료")
}

func run() {
	// 로그 메서드를 쓰면 파일에 출력됨
	log.Print("테스트")
}
