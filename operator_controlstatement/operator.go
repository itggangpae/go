package main

import "fmt"

func main() {
	a := 1
	a-- //0:정수1을 1 감소시켜서0
	fmt.Printf("a-- : %d\n", a)

	a = 8 + 10/3
	fmt.Println(a) // 11

	a = (8 + 10) / 3
	fmt.Println(a) // 6
}
