package main

import "fmt"

func main() {
	i := 10
	if i >= 5 {
		fmt.Println("5 이상")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
