package main

import "fmt"

func printSqure(a int) {
	a *= a
	fmt.Println(a)
}

func main() {
	//지역 변수 선언
	a := 3
	//함수 호출 - 값에 의한 전달
	printSqure(a)
	//원본은 변경되지 않음
	fmt.Println(a)
}
