package main

import (
	"fmt"
)

func main() {
	ar := [3]int{10, 20, 30}
	//배열의 데이터 개수
	n := len(ar)
	fmt.Println("데이터 개수:", n)

	//배열의 순회
	for i := 0; i < n; i++ {
		fmt.Printf("[%d]:%d\n", i, ar[i])
	}

	// 하나의 변수로 받으면 인덱스가 들어감
	for index := range ar {
		fmt.Println(index)
	}

	// value에는 값 대신 인덱스가 들어감
	for index, value := range ar {
		fmt.Println(index, ":", value)
	}
}
