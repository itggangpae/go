package main

import (
	"fmt"
)

func main() {
	//배열 생성
	var myAr = [5]int{1, 2, 3, 4, 5}
	//배열의 일부분 잘라내기 - 참조 복사
	var subAr = myAr[2:4]
	myAr[2] = 30
	fmt.Println(myAr)
	fmt.Println(subAr)

	//슬라이스 생성
	var mySlice = []int{1, 2, 3, 4, 5}
	//슬라이스 잘라내기 - 참조 복사
	var subSlice = mySlice[2:4]
	mySlice[2] = 30
	fmt.Println(mySlice)
	fmt.Println(subSlice)

	// 길이가 큰 슬라이스
	var myBigSlice = []int{1, 2, 3, 4, 5, 6}
	// 작은 슬라이스를 생성
	var mySubSlice = make([]int, 2)
	// myBigSlice에서 mySubSlice로 2개 요소의 값을 복사
	copy(mySubSlice, myBigSlice[2:4])
	myBigSlice[2] = 30
	fmt.Println(myBigSlice)
	fmt.Println(mySubSlice)
}
