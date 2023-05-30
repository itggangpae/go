package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("컴파일러는 ", runtime.Compiler, "이고")
	fmt.Println("현재 머신은", runtime.GOARCH, "머신")
	fmt.Println("Go version:", runtime.Version())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
