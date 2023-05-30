package main

import (
	"fmt"
	"unsafe"
)

type User1 struct {
	Age   int     // 8바이트
	Score float64 // 8바이트
}

type User2 struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

func main() {
	user1 := User1{23, 77.2}
	fmt.Println(unsafe.Sizeof(user1))

	user2 := User2{23, 77.2}
	fmt.Println(unsafe.Sizeof(user2))
}
