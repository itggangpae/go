package main

// 피보나치 수열을 재귀 호출로 구현한 함수
func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci1(n-1) + fibonacci1(n-2) //재귀 호출
}

// 피보나치 수열을 반복문으로 구현한 함수
func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	one := 1
	two := 0
	rst := 0
	for i := 2; i <= n; i++ { //반복문
		rst = one + two
		two = one
		one = rst
	}
	return rst
}
