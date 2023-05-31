package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("작업 시작\n")

	wg.Add(3)
	go MakeBody(tireCh) //Go 루틴 생성
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("작업 종료")
}

func MakeBody(tireCh chan *Car) { //차체 생산
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			//자동차 생산 시작
			car := &Car{}
			//차체 생산
			car.Body = "Sports car"
			//바퀴 작업에 시그널 전송
			tireCh <- car
		case <-after: //10초 뒤 종료
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) { //바퀴 설치
	for car := range tireCh {
		time.Sleep(time.Second)
		//바퀴 생성
		car.Tire = "Winter tire"
		//도색 작업에 시그널 전송
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) { //도색
	for car := range paintCh {
		// Make a body
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime) //경과 시간 출력
		fmt.Printf("%.2f 자동차 완성: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
