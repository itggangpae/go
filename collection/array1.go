package main

import (
	"fmt"
)

func main() {
	//zero value를 갖는 배열 생성
	var ar [3]int
	//초기화를 이용한 배열 생성
	ar = [3]int{1, 2, 3}
	//배열 전체 출력
	fmt.Println(ar)
}
