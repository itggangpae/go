package main

import (
	"fmt"
)

func main() {
	ar := [5]int{10, 20, 30, 40, 50}
	fmt.Println(ar[2:4])
	fmt.Println(ar[:4])
	fmt.Println(ar[2:])
}
