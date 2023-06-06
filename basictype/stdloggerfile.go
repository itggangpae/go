package main

import (
	"log"
	"os"
	"time"
)

func main() {
	fpLog, err := os.OpenFile(time.Now().Local().Format("2006-01-02"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	// 표준로거를 파일 로그로 변경
	log.SetOutput(fpLog)

	run()
	log.Println("프로그램 종료")
}

func run() {
	// 로그 메서드를 쓰면 파일에 출력됨
	log.Print("테스트")
}
