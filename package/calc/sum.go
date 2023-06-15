package calc

import "fmt"

var globalVal = 0

func init() {
	fmt.Println("init()", globalVal)
	globalVal++
}

func Sum(a int, b int) int {
	return a + b
}
