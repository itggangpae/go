package main

import (
	"fmt"
	"os"
)

func Helloworld() {
	file, err := os.Open("test.txt")
	//defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1024)

	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
	file.Close()
}

func main() {
	Helloworld()
	fmt.Println("Done")
}
