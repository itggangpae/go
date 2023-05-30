package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("test_1.txt")
	defer file1.Close()
	fmt.Fprintln(file1, 1, 1.1, "Adam")

	file2, _ := os.Create("test_2.txt")
	defer file2.Close()
	fmt.Fprintf(file2, "%d,%f,%s", 1, 1.1, "Adam")
}
