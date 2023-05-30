package main

import "fmt"

type MyInterface interface {
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

// 인터페이스를 매개변수로 받는 함수
func printNameAndAge(i MyInterface) {
	fmt.Println(i.GetName(), i.GetAge())
}

func main() {
	p := Person{"Adam", 51}
	printNameAndAge(p)
}
