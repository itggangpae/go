package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) method() {
	fmt.Println("Super 의 Method")
}

type Student struct {
	Person
	school string
	grade  int
}

func (s *Student) method() {
	//상위 구조체의 메서드 호출
	s.Person.method()
	fmt.Println("Sub 의 Method")
}

func main() {
	var s Student
	//상위 구조체의 메서드 호출
	s.Person.method()
	//하위 구조체의 재정의된 메서드 호출
	s.method()
}
