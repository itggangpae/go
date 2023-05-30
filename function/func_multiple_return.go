package main

import "fmt"

func SumAndDiff1(a int, b int) (int, int) { // int형 리턴값이 두 개인 함수 정의
	return a + b, a - b // 리턴할 값 또는 변수를 차례대로 나열. 합과 차를 동시에 리턴
}

func SumAndDiff2(a int, b int) (int, int) {
	return a + b, a - b // 리턴할 값 또는 변수를 차례대로 나열
}

func hello() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5 // 리턴할 값을 차례대로 나열
}

func SumAndDiff3(a int, b int) (sum int, diff int) { // 리턴값을 각각 sum, diff로 이름을 정함
	sum = a + b  // 리턴값 변수 sum에 합 대입
	diff = a - b // 리턴값 변수 diff에 차 대입
	return
}

func main() {
	sum, diff := SumAndDiff1(6, 2) // 변수 두 개에 리턴값 두 개를 대입
	fmt.Println(sum, diff)         // 8 4: 합과 차

	_, diff = SumAndDiff2(6, 2) // 첫 번째 리턴값은 _으로 생략, 두 번째 리턴값만 사용
	fmt.Println(diff)           // 4: 차

	a, _, c, _, e := hello() // 2, 4번째 리턴값은 생략
	fmt.Println(a, c, e)     // 1 3 5

	sum, diff = SumAndDiff3(6, 2)
	fmt.Println(sum, diff) // 8 4
}
