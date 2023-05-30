package main

import (
	"fmt"
)

func S1() {
	fmt.Println("In S1()")
	return

	fmt.Println("Leaving S1()")
}

func S2() {
	return
	fmt.Println("Hello!")
}
