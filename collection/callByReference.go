package main

import "fmt"

func callByValue(n int) {
	n = 2
	fmt.Println(n)
}

func callByReference(n *int) {
	*n = 2
	fmt.Println(*n)
}

func main() {
	var n int = 1
	callByValue(n)
	fmt.Println(n)

	callByReference(&n)
	fmt.Println(n)
}
