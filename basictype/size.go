package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var sum float64 = 0.0
	fmt.Println("메모리 크기:", unsafe.Sizeof(sum))
}
