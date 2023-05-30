package main

import "fmt"

func main() {
	var r1 rune = '한'
	var r2 rune = '\ud55c'     // 한
	var r3 rune = '\U0000d55c' // 한

	fmt.Println(r1, r2, r3)

	/*
		var r1 rune = "한"   // 컴파일 에러
		var r2 rune = '한글' // 컴파일 에러
		var r3 rune = "한글" // 컴파일 에러
	*/
}
