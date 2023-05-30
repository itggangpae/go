package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 string = "Hello, world!\n"
	var s2 string = "안녕하세요\n"
	var s3 string = "\ud55c\uae00"             // 한글:유니코드 문자로 저장
	var s4 string = "\U0000d55c\U0000ae00"     // 한글: 유니코드 문자 코드로 저장
	var s5 string = "\xed\x95\x9c\xea\xb8\x80" // 한글: UTF-8 바이트 값으로 저장

	var s7 string = `안녕하세요
	Hello, world!`

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s7)

	//문자 단위 접근
	aString := "Hello World!"
	fmt.Println("First character:", string(aString[0]))

	//rune 단위 접근
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	//문자 단위 접근
	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	var str1 string = "한글"
	var str2 string = "Hello"
	fmt.Println(len(str1))                    // 6: UTF-8 인코딩의 바이트 길이이므로 6
	fmt.Println(len(str2))                    // 5: 알파벳 5글자이므로 5
	fmt.Println(utf8.RuneCountInString(str1)) // 2: 문자열의 실제 길이를 구함

}
