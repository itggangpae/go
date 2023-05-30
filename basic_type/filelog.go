package main

import (
	"log"
	"os"
)

var myLogger *log.Logger

func main() {
	// 로그파일 오픈
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	myLogger = log.New(fpLog, "Debug: ", log.Ldate|log.Ltime|log.Lshortfile)

	//....
	run()

	myLogger.Println("프로그램 종료")
}

func run() {
	myLogger.Print("테스트")
}
