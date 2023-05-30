package main

import (
	"fmt"
)

func main() {
	ar := [3]int{10, 20, 30}
	//배열의 데이터 개수
	n := len(ar)
	fmt.Println(n)

	//배열의 순회
	for i := 0; i < n; i++ {
		fmt.Println(ar[i])
	}

	// value에는 값 대신 인덱스가 들어감
	for value := range ar {
		fmt.Println(value)
	}

	// value에는 값 대신 인덱스가 들어감
	for _, value := range ar {
		fmt.Println(value)
	}

}
