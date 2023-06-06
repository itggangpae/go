package main

//사용할 패키지 import
import (
	"fmt"
)

// 전역 변수 선언
var Global int = 33
var AnotherGlobal = 12.345678

func main() {
	//지역 변수 선언
	i := 33 //초기값을 부여해서 생성
	j := 33 //초기값 없이 생성

	fmt.Println("j의 초기값:", j)
	j = Global

	//서식을 설정해서 출력
	fmt.Printf("전역변수\n Global=%d\t AnotherGlobal=%.3f\n지역변수\n i=%d\t j=%05d\n", Global, AnotherGlobal, i, j)
}
