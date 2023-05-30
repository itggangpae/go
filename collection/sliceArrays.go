package main

import (
	"fmt"
)

func change(s []string) {
	s[0] = "Change_function"
}

func main() {
	//배열 생성
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	//슬라이스 생성
	var S0 = a[0:1]
	fmt.Println(S0)
	S0[0] = "S0"

	var S12 = a[1:3]
	fmt.Println(S12)

	S12[0] = "S12_0"
	S12[1] = "S12_1"

	//슬라이스를 이용했으므로 원본의 데이터도 같이 변경
	fmt.Println("a:", a)

	//슬라이스를 배열로 변환
	change(S12)
	fmt.Println("a:", a)

	// 용량 확인
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))

	// 데이터 추가
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	a[0] = "-N1"

	// 데이터 추가 - 슬라이스의 용량이 배열의 용량보다 커 짐
	S0 = append(S0, "N4")

	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
	// 연결이 끊어짐
	a[0] = "-N1-"
	a[1] = "-N2-"

	fmt.Println("S0:", S0)
	fmt.Println("a: ", a)
	fmt.Println("S12:", S12)
}
