package main

import "fmt"

type MyInter interface {
	GetName() string
	GetAge() int
}

type Person struct {
	name string
	age  int
}

func (p Person) GetName() string {
	return p.name
}
func (p Person) GetAge() int {
	return p.age
}

func main() {
	var myInterfaceValue MyInter

	var p = Person{}
	p.name = "아담"
	p.age = 51

	// 인터페이스 사용
	myInterfaceValue = p
	fmt.Println(myInterfaceValue.GetName())
	fmt.Println(myInterfaceValue.GetAge())
}
