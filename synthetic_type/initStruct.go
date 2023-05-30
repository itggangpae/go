package main

import (
	"fmt"
)

type Entry struct {
	Name    string
	Surname string
	Year    int
}

func main() {
	var entry1 = Entry{"adam", "itstudy", 1997}
	var entry2 = Entry{}
	var entry3 *Entry = new(Entry)

	fmt.Println(entry1)
	fmt.Println(entry2)
	fmt.Println(entry3) //구조체 포인터 이므로 앞에 & 가 붙음

	entry2.Name = "suzi"
	entry2.Surname = "bae"
	entry2.Year = 1994
	fmt.Println(entry2)

	entry3.Name = "jieun" //구조체 포인터에 접근할 때도 .을 사용
	entry3.Surname = "lee"
	entry3.Year = 1993
	fmt.Println(entry3)

}
