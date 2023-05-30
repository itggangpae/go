package main

import "fmt"

// int형 매개변수 a, b 그리고 int 형 리턴값을 가지는 함수 정의
func sum1(a int, b int) int {
	return a + b
}

// 리턴 값 변수의 이름을 r로 지정
func sum2(a int, b int) (r int) {
	r = a + b // 리턴값 변수 r에 값 대입
	return    // 리턴값 변수를 사용할 때는 return 뒤에 변수를 지정하지 않음
}

func main() {
	r := sum1(1, 2)
	fmt.Println(r) // 3

	r = sum2(1, 2)
	fmt.Println(r)
}
