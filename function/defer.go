package main

import "fmt"

func func1() {
	fmt.Println("Hello")
}

func func2() {
	fmt.Println("Go")
}

func main() {
	//함수의 호출이 지연
	defer func1() // 현재 함수(main())가 끝나기 직전에 호출
	func2()

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
