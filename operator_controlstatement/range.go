package main

import (
	"fmt"
)

func main() {
	//0부터 5까지
	for i := range [5]int{} {
		fmt.Println(i)
	}
}
