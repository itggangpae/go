package main

import "fmt"

// 함수의 타입에 새로운 이름 부여
type funcType func(int, int) int

// 함수를 매개변수로 받는 함수
func calc1(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// 새로 만든 타입으로 함수를 설정
func calc2(f funcType, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	add := func(a int, b int) int {
		return a + b
	}

	result1 := calc1(add, 10, 20)
	fmt.Println(result1)

	result2 := calc2(func(a int, b int) int { return a + b }, 10, 20)
	fmt.Println(result2)
}
