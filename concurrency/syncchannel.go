package main

import (
	"fmt"
	"time"
)

func runLoopSend(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
}

/*
func runLoopReceive(ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received value:", i)
	}
}
*/

func runLoopReceive(ch chan int) {
	for i := range ch {
		fmt.Println("Received value:", i)
	}
}

func main() {
	//채널 생성
	myChannel := make(chan int)
	go runLoopSend(10, myChannel)
	go runLoopReceive(myChannel)
	time.Sleep(10 * time.Second)
}
