package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//디바이스의 총 CPU 개수를 반환하고 그 값을 CPU 사용 값으로 설정
	fmt.Println(runtime.GOMAXPROCS(0))
	// 현재 설정값 출력, 1미만이기 때문에 설정값 바꾸지 않음
	var wait sync.WaitGroup
	wait.Add(100)

	for i := 0; i < 100; i++ {
		go func(n int) {
			defer wait.Done()
			fmt.Println(n)
		}(i)
	}

	wait.Wait()
}
