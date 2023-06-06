package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("현재 시간:", start)

	dateString := "25 December 2021 14:10"
	//날짜만 존재하는지 확인
	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	}

	//날짜 와 시간인지 확인
	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// 숫자 형식으로 날짜 외 시간이 표현
	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// 시간만 존재
	d, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}
}
