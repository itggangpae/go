package main

import "fmt"

// 함수를 리턴
func outer() func() {
	a := 1 // 지역 변수는 함수가 호출되서 실행이 된 후 소멸되지만
	return func() {
		fmt.Printf("a = %d\n", a)
		a++
		//내부 함수이므로 함수를 호출 할 때마다 변수 a의 값을 사용할 수 있음
	}
	//익명 함수를 리턴
}

func main() {
	f := outer() // outer 함수를 실행하여 리턴값으로 나온 함수를 변수에 저장

	f()
	f()
}
