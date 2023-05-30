package main

import (
	"fmt"
	"time"
)

func runSomeLoop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Printing:", i)
	}
}
func main() {
	//runSomeLoop(10)

	go runSomeLoop(10)
	// main 고루틴 2초 동안 대기
	time.Sleep(2 * time.Second)

	fmt.Println("Hello, playground")
}
