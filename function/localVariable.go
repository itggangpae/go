package main

import "fmt"

func exampleFunc1() {
	var a int = 10 //지역변수 선언

	a++

	fmt.Println("exampleFunc1의 a는 ", a)
}

func exampleFunc2() {
	var b int = 20 //지역변수 선언
	var c int = 30 //지역변수 선언

	b++
	c++

	fmt.Println("b와 c는 ", b, c)
}

func main() {
	var a int = 28 //지역변수 선언

	exampleFunc1()
	exampleFunc2()

	fmt.Println("main의 a는", a)
}
