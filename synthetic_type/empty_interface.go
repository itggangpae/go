package main

import (
	"fmt"
	"strconv"
)

// 빈 인터페이스를 사용하여 모든 타입을 받음
func formatString(arg interface{}) string {
	//인터페이스에 저장된 타입에 따라 case 실행
	switch arg.(type) {
	case int: // arg가 int형이라면
		i := arg.(int)         // int형으로 값을 가져옴
		return strconv.Itoa(i) // strconv.Itoa 함수를 사용하여 i를 문자열로 변환
	case float32: // arg가 float32형이라면
		f := arg.(float32) // float32형으로 값을 가져옴
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
		// strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case float64: // arg가 float64형이라면
		f := arg.(float64) // float64형으로 값을 가져옴
		return strconv.FormatFloat(f, 'f', -1, 64)
		// strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case string: // arg가 string이라면
		s := arg.(string) // string으로 값을 가져옴
		return s          // string이므로 그대로 리턴
	case Person: // arg의 타입이 Person이라면
		p := arg.(Person)                         // Person 타입으로 값을 가져옴
		return p.name + " " + strconv.Itoa(p.age) // 각 필드를 합쳐서 리턴
	case *Person: // arg의 타입이 Person 포인터라면
		p := arg.(*Person)                        // *Person 타입으로 값을 가져옴
		return p.name + " " + strconv.Itoa(p.age) // 각 필드를 합쳐서 리턴
	default:
		return "Error"
	}
}

type Person struct { // Person 구조체 정의
	name string
	age  int
}

func main() {
	fmt.Println(formatString(3))
	fmt.Println(formatString(3.14))
	fmt.Println(formatString("Hello GO!"))

	fmt.Println(formatString(Person{"Adam", 51}))
	fmt.Println(formatString(&Person{"Jessica", 49}))

	var hunt = new(Person)
	hunt.name = "Hunt"
	hunt.age = 9

	fmt.Println(formatString(hunt))
}
