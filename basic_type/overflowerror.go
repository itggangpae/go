package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 uint8 = math.MaxUint8 + 1 // 오버플로우 컴파일 에러 발생
	var num2 uint16 = 0 - 1            // 오버플로우 컴파일 에러 발생
	fmt.Println(num1)
	fmt.Println(num2)
}
