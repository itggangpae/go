package main

//사용할 패키지 import
import (
	"fmt"
	"math"
)

// 전역 변수 선언
var Global int = 1234
var AnotherGlobal = -5678

func main() {
	//지역 변수 선언
	var j int                   //초기값 없이 생성
	i := Global + AnotherGlobal //초기값을 부여해서 생성
	fmt.Println("j의 초기값 - Zero Value:", j)
	j = Global
	// math.Abs()의 매개변수는 float64 타입이므로 적절하게 타입을 변환
	k := math.Abs(float64(AnotherGlobal))
	//서식을 설정해서 출력
	fmt.Printf("Global=%d, i=%d, j=%d k=%.2f.\n", Global, i, j, k)
}
