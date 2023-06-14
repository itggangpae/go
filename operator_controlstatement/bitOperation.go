package main

import "fmt"

func main() {
	num1 := 15 //00001111
	num2 := 20 //00010100

	fmt.Printf("num1 & num2 : %08b, %d\n", num1&num2, num1&num2)
	fmt.Printf("num1 | num2 : %08b, %d\n", num1|num2, num1|num2)
	fmt.Printf("num1 ^ num2 : %08b, %d\n", num1^num2, num1^num2)
	fmt.Printf("num1 &^ num2 : %08b, %d\n", num1&^num2, num1&^num2)

	fmt.Printf("num1 << 4 : %08b, %d\n", num1<<4, num1<<4)
	fmt.Printf("num2 >> 2 : %08b, %d\n", num2>>2, num2>>2)
}
