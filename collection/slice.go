package main

import (
	"fmt"
)

func main() {
	//배열 생성
	var myAr = [5]int{1, 2, 3, 4, 5}
	//배열의 일부분 잘라내기 - 참조 복사
	var subAr = myAr[2:4]
	fmt.Println("subAr:", subAr)

	//원본 데이터 변경
	myAr[2] = 30
	//슬라이스가 가리키는 데이터도 변경
	fmt.Println("subAr:", subAr)

	//슬라이스 생성
	var mySlice = []int{1, 2, 3, 4, 5}
	//슬라이스 잘라내기 - 참조 복사
	var subSlice = mySlice[2:4]
	fmt.Println("subSlice:", subSlice)

	mySlice[2] = 30
	fmt.Println("subSlice:", subSlice)

	//슬라이스를 생성
	var makeSlice = make([]int, 2, 4)
	fmt.Println("makeSlice:", makeSlice)
	fmt.Println("데이터 개수:", len(makeSlice))
	fmt.Println("슬라이스 용량:", cap(makeSlice))
}
