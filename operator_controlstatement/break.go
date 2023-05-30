package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		if i > 2 {
			break
		}
		i++
		fmt.Println(i)
	}
}
