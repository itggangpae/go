package main

import "fmt"

func main() {
	var slice []int //0개 cap을 갖는 슬라이스

	slice, allocCnt := append1000times(slice)
	fmt.Println("allocCnt:", allocCnt, "cap:", cap(slice))

	var slice2 = make([]int, 0, 1000) //1000개 cap을 갖는 슬라이스
	slice, allocCnt = append1000times(slice2)
	fmt.Println("allocCnt:", allocCnt, "cap:", cap(slice))
}

func append1000times(slice []int) ([]int, int) {
	var lastCap int = cap(slice)
	var allocCnt int = 0
	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
		if lastCap != cap(slice) { //capacity 가 바뀔때 allocCnt 증가
			allocCnt++
			lastCap = cap(slice)
		}
	}
	return slice, allocCnt
}
