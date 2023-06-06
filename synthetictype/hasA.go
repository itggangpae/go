package main

import "fmt"

// 구조체 생성
type Person struct {
	name string
	age  int
}

// Person 구조체의 메서드 생성
func (p *Person) hasA() {
	fmt.Println("Has A")
}

// 다른 구조체 인스턴스를 포함하는 구조체
type Student struct {
	p      Person
	school string
	grade  int
}

func main() {
	var s Student
	s.p.hasA()
}
