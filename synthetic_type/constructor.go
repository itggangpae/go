package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height} // 구조체 인스턴스를 생성한 뒤 포인터를 리턴
}

func main() {
	rect := NewRectangle(20, 10)
	fmt.Println(rect) // &{20 10}

	rect_ptr := &Rectangle{20, 10} // 구조체를 초기화한 뒤 메모리 주소를 대입
	fmt.Println(rect_ptr)          // &{20 10}
}
