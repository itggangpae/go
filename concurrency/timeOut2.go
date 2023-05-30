package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var result = make(chan bool)

func timeout(t time.Duration) {
	temp := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)
	}()

	select {
	case <-temp:
		result <- false
	case <-time.After(t):
		result <- true
	}
}

func main() {
	arguments := os.Args

	if len(arguments) != 2 {
		fmt.Println("밀리초 단위의 타임아수 시간 설정")
		return
	}

	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout 간격은 %s\n", duration)

	go timeout(duration)

	val := <-result
	if val {
		fmt.Println("Time out!")
	} else {
		fmt.Println("OK")
	}
}
