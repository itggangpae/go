package main

import (
	"fmt"
	"strconv"
)

// 구조체 선언
type record struct {
	Field1 int
	Field2 string
}

func main() {
	//구조체 슬라이스 생성
	S := []record{}

	//슬라이스에 데이터 추가
	for i := 0; i < 10; i++ {
		text := "text" + strconv.Itoa(i)
		temp := record{Field1: i, Field2: text}
		S = append(S, temp)
	}

	//첫번째 데이터의 필드 출력
	fmt.Println("Index 0:", S[0].Field1, S[0].Field2)
	//슬라이스 길이 출력
	fmt.Println("Number of structures:", len(S))
	//첫번째 필드의 합계
	sum := 0
	for _, k := range S {
		sum += k.Field1
	}
	fmt.Println("Sum:", sum)
}
