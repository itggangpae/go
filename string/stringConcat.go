package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "Hello"
	str1 += "World"
	fmt.Println(str1)

	str1 = "Hello"
	var builder strings.Builder
	builder.WriteString("Hello" + "World")
	fmt.Println(builder.String())
}
