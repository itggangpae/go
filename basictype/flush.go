package main

import (
	//"bufio" // io를 담당하는 패키지
	"fmt"
	//"os" // 표준 입출력 등을 가지고 있는 패키지
)

func main() {
	//stdin := bufio.NewReader(os.Stdin) // 표준 입력을 읽는 객체

	var a int
	var b int

	fmt.Print("정수 2개를 입력:")
	n, err := fmt.Scanln(&a, &b)
	if err != nil { // 에러 발생 시
		fmt.Println(err) // 에러 출력
		//stdin.ReadString('\n') // 표준 입력 스트림 지우기
	} else {
		fmt.Println(n, a, b)
	}

	fmt.Print("정수 2개를 입력:")
	n, err = fmt.Scanln(&a, &b) // 다시 입력받기
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
