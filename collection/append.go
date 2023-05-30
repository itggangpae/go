package main

import (
	"fmt"
)

func main() {
	ar := []int{1, 2, 3}
	ar = append(ar, 4)
	fmt.Println(ar)

	br := []int{5, 6, 7}
	ar = append(ar, br...)
	fmt.Println(ar)

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a), cap(a))
	a = append(a, 6, 7)
	fmt.Println(len(a), cap(a))
}
