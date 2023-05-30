package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 10.0
	for i := 0; i < 10; i++ {
		a = a - 0.1
	}
	fmt.Println(a)        // 9.000000000000004: 반올림 오차 발생
	fmt.Println(a == 9.0) // false: 실수는 ==로 비교하면 안 됨

	const epsilon = 1e-14                   // Go 언어 머신 엡실론
	fmt.Println(math.Abs(a-9.0) <= epsilon) // true: 연산한 값과 비교할 값의 차이를 구한 뒤 머신 엡실론보다 작거나 같은지
}
