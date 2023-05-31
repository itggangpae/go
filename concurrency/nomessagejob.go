package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            //1초 간격 시그널
	terminate := time.After(10 * time.Second) //10초 이후 시그널

	for {
		select { //tick, terminate, ch 순서로 처리
		case <-tick:
			fmt.Println("일정한 간격의 작업")
		case <-terminate:
			fmt.Println("작업 종료")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
