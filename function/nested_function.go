package main

import "fmt"

// 함수 안에서
func main() {
	//내부 함수는 익명 함수 형태로 선언 및 정의
	diff := func(a, b int) int {
		return a - b
	}

	r := diff(1, 2) //내부 함수 호출

	fmt.Println(r) // -1
}
