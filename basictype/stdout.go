package main

import (
	"log"
	"os"
)

var looger *log.Logger

func main() {
	looger = log.New(os.Stdout, "Debug: ", log.LstdFlags)

	run()

	looger.Println("프로그램 종료")
}

func run() {
	looger.Print("테스트")
}
