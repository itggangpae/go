package main

import (
	"fmt"
	"time"
)

func main() {
	// 5개 이상의 정수를 저장할 수 있는 버퍼 채널 생성
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		// number에 대입
		case numbers <- i * i:
			fmt.Println("작업 수행:", i)
			time.Sleep(1000)
		//버퍼가 다 차면 수행
		default:
			fmt.Println("더 이상 저장할 수 없음:  ", i, " ")
		}
	}
	fmt.Println()

	for {
		select {
		case num := <-numbers:
			fmt.Print("*", num, " ")
		default:
			fmt.Println("더 이상 데이터를 읽을 수 없음")
			return
		}
	}
}
