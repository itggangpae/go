package main

import (
	"encoding/xml"
	"fmt"
)

// 이름 과 타입에 관한 정보가 추가로 제공되어야 함
// xml.Name 은 레코드의 이름 이고 나머지는 태그
type Employee struct {
	XMLName   xml.Name `xml:"employee"`
	ID        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Height    float32  `xml:"height,omitempty"`
	Address
	Comment string `xml:",comment"`
}

type Address struct {
	City, Country string
}

func main() {
	r := Employee{ID: 7, FirstName: "Mihalis", LastName: "Tsoukalos"}
	r.Comment = "Technical Writer + DevOps"
	r.Address = Address{"SomeWhere 12", "12312, Greece"}

	output, err := xml.MarshalIndent(&r, "  ", "    ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	output = []byte(xml.Header + string(output))
	fmt.Printf("%s\n", output)
}