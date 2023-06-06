package main

import "fmt"

func main() {
	str := "안녕하세요 문자엽니다."

	fmt.Println("문자열의 내용:", str)
	fmt.Println("문자열의 메모리 크기:", len(str))
	fmt.Println("문자열의 글자 수:", len([]rune(str)))
}
