/*
package main

import (
	"fmt"
	"time"
)

func f(n int) {
	time.Sleep(time.Duration(10)) // 랜덤한 시간 동안 대기
	fmt.Println(n)                // n 출력
}

func main() {
	for i := 0; i < 100; i++ { // 100번 반복하여
		go f(i) // 고루틴 100개 생성
	}
	//콘솔 입력 대기
	fmt.Scanln()
}
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func f(n int) {
	time.Sleep(time.Duration(100)) // 랜덤한 시간 동안 대기
	fmt.Println(n)                 // n 출력
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

	fmt.Println(runtime.GOMAXPROCS(0)) // 설정 값 출력
	fmt.Println("======================")
	for i := 0; i < 100; i++ {
		go f(i)
	}

	fmt.Scanln()
}
