package coverage

import "fmt"

// 도달할 수 없는 코드를 가진 함수
func f1() {
	if true {
		fmt.Println("무조건 호출!")
	} else {
		fmt.Println("절대로 호출될 수 없는 코드!") //도달 할 수 없는 코드
	}
}

// 음수를 대입하면 동작하지 않고 양수를 대입하면 첫번째 코드만 수행되는 함수
func f2(n int) int {
	if n >= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return f2(n-1) + f2(n-2)
	}
}
