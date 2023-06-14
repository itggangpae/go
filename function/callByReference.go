package main

import "fmt"

func printSqure(a *int) {
	*a *= *a
	fmt.Println(*a)
}
func main() {
	//지역 변수 선언
	a := 3
	//참조를 위한 a의 참조를 매개변수로 전달
	printSqure(&a)
	//데이터가 수정될 수 도 있음
	fmt.Println(a)
}
