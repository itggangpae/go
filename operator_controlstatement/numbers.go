package main

import (
	"fmt"
)

func main() {
	c1 := 12 + 1i
	c2 := complex(5, 7)
	//자료형 확인
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)

	//값과 자료형 확인
	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Type of c3: %T\n", c3)

	cZero := c3 - c3
	fmt.Println("cZero:", cZero)

	//정수 자료형 사용
	x := 12
	k := 5
	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)

	//정수 연산의 결과는 정수
	div := x / k
	fmt.Println("div", div)

	//실수 연산
	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)

	y := 4 / 2.3
	fmt.Println("y:", y)

	divFloat := float64(x) / float64(k)
	fmt.Println("divFloat", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)
}
