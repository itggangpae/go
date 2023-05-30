package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 float32
	var s string

	file1, _ := os.Open("test_1.txt")
	defer file1.Close()
	fmt.Fscanln(file1, &num1, &num2, &s)
	fmt.Println(num1, num2, s)

	file2, _ := os.Open("test_2.txt")
	defer file2.Close()
	fmt.Fscanf(file2, "%d,%f,%s", &num1, &num2, &s)
	fmt.Println(num1, num2, s)
}
