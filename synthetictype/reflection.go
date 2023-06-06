package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

// 다른 구조체를 포함하는 구조체 선언
type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"adam", "1234567"}}

	r := reflect.ValueOf(A) //A 의 타입의 값을 리턴
	fmt.Println("String value:", r.String())

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType) //A의 타입 리턴
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)

	//내부 종류의 타입을 리턴
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())

		//포함하고 있는 타입 체크
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		// 구조체 인지 확인
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}

		// 타입 확인
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
