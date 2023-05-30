package main

import "fmt"

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)
	// Pointer to f
	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)

	var numPtr *int = new(int) // new 함수로 공간 할당
	*numPtr = 1                // 역참조로 포인터형 변수에 값을 대입
	fmt.Println(*numPtr)       // 1: 포인터형 변수에서 값을 가져오기
}
