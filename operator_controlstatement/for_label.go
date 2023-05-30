package main

import (
	"fmt"
)

func main() {
Loop: // Loop 레이블을 지정
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("Label Use")
}
