package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 처리할 요청을 관리하는 구조체
type Client struct {
	id      int
	integer int
}

// Client 의 데이터와 이로부터 발생한 결과를 저장하는데 사용하는 구조체
type Result struct {
	job    Client
	square int
}

// 버퍼 채널 생성
var size = runtime.GOMAXPROCS(0)
var clients = make(chan Client, size)
var data = make(chan Result, size)

// clients 채널의 데이터를 읽어서 요청을 처리하는 함수
// 처리가 끝난 후에는 data 채널에 결과를 기록
func worker(wg *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Result{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

// 요청을 생성한 후 clients 버퍼 채널로 요청을 전달
func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		return
	}
	//Command Line의 매개변수를 읽어서 작업 과 워커의 개수를 설정
	//워커의 수가 clients 버퍼 채널의 크기보다 크면 생성되는 고루틴의 개수는 clients 채널의 크기와 같아지고
	//작업의 개수가 워커의 개수보다 크다면 작은 작은 단위로 나누어서 처리
	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	//요청 생성
	go create(nJobs)

	//프로그램 실행을 블록 시키기 위해서 호출
	finished := make(chan interface{})

	//finished <- true 구문은 for range 루프가 끝난 뒤 프로그램이 블록 상태가 해제 된다는 것을 의미
	//wg.WaitQ가 반환된 다음 data 채널이 닫히고 for range 루프가 종료되며 모든 워커의 작업이 끝났다는 것을 의미
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()

	//요청을 처리할 worker 고루틴을 필요한 만큼 생성
	var wg sync.WaitGroup
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(data)

	//채널이 닫히기 전까지 블록
	fmt.Printf("Finished: %v\n", <-finished)
}
