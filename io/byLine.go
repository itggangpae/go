package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	//읽기 변수 생성
	r := bufio.NewReader(f)
	for {
		//읽은 문자열 과 에러 변수 반환
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byWord <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
