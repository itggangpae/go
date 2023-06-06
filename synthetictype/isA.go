package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) isA() {
	fmt.Println("Is A")
}

func (p *Person) hello() {
	fmt.Println("Hello")
}

type Student struct {
	Person
	school string
	grade  int
}

func main() {
	var s Student
	s.Person.isA()
	s.isA()
}
