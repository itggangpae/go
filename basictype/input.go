package main

import (
	"fmt"
)

func main() {
	fmt.Print("이름을 입력하세요: ")
	var name string
	//키보드 입력을 대기 한 후 입력 후 Enter를 누르면 name에 입력된 내용을 대입
	fmt.Scanln(&name)
	fmt.Println("당신의 이름은 ", name)
}
