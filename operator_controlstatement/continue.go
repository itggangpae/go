package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		if i%3 == 2 {
			continue
		}
		fmt.Println(i)
	}
}
