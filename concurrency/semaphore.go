package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/sync/semaphore"
)

// 동시에 수행되는 고루틴의 개수 설정
var Workers = 4
var sem = semaphore.NewWeighted(int64(Workers))

// 고루틴으로 수행되는 함수
func worker(n int) int {
	square := n * n
	time.Sleep(5 * time.Second)
	fmt.Println("실행")
	return square
}

func main() {
	//Command Line 인수에 실행할 고루틴 개수가 없으면 종료
	if len(os.Args) != 2 {
		fmt.Println("작업의 개수가 필요합니다.")
		return
	}
	//두번째 매개변수를 정수로 변환
	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	//결과를 저장하기 위한 슬라이스 생성
	var results = make([]int, nJobs)

	//데이터 저장을 위한 context 생성
	ctx := context.TODO()

	for i := range results {
		//Lock을 취득
		err = sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Println("세마포어를 사용할 수 없음:", err)
			break
		}
		//고루틴 실행
		go func(i int) {
			defer sem.Release(1)
			temp := worker(i)
			results[i] = temp
		}(i)
	}

	//세마포어를 생성해보고 에러가 발생하면 메시지 출력
	err = sem.Acquire(ctx, int64(Workers))
	if err != nil {
		fmt.Println(err)
	}

	//결과 출력
	for k, v := range results {
		fmt.Println(k, "->", v)
	}
}
