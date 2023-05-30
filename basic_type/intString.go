package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "123"

	//문자열을 정수로 변환
	n, err := strconv.Atoi(str)
	//에러가 발생하면 메러 메시지 출력
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("n + 1 = %d of Type %T\n", n+1, n)

	//정수를 문자열로 변환
	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)
}
