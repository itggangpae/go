package main

import "fmt"

func main() {
	//변수 생성
	var b1 bool = true
	var b2 bool = false

	fmt.Println(b1)
	fmt.Println(b2)

	//논리 연산
	fmt.Println(true && true)   // true
	fmt.Println(true && false)  // false
	fmt.Println(false && false) // false
	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(false || false) // false
	fmt.Println(!true)          // false
	fmt.Println(!false)         // true

	//비교 연산
	var num1 int = 3
	var num2 int = 10

	fmt.Println(num1 > num2)  // false
	fmt.Println(num1 < num2)  // true
	fmt.Println(num1 != num2) // true
	fmt.Println(num1 >= num2) // false
	fmt.Println(num1 <= num2) // true
}
