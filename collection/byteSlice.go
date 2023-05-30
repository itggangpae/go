package main

import "fmt"

func main() {
	//바이트 슬라이스
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)
	b = []byte("Byte slice P")
	fmt.Println("Byte slice:", b)

	//문자열로 출력
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))

	//슬라이스 길이
	fmt.Println("Length of b:", len(b))
}
