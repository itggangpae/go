package main

import "fmt"

func main() {
	str := "H 드" //한영이 섞인 문자열

	fmt.Println("바이트 단위 접근")
	//인덱스 단위 접근
	for i := 0; i < len(str); i++ { //문자열 크기를 얻어 순회
		//바이트 단위로 출력
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}
	fmt.Println("룬 단위 접근")
	//룬 단위 접근
	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}
	fmt.Println("range를 이용한 접근")
	//range 이용
	for _, v := range str {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", v, v, v)
	}
}
