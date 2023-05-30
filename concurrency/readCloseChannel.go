package main

import "fmt"

func main() {
	//willClose := make(chan complex64) //에러 - 채널에 2개의 데이터를 쓸려고 해서 읽고 난 후 비우고 기록해야 함
	willClose := make(chan complex64, 3) //Zero Value를 읽음

	// 채널에 기록
	willClose <- -1
	willClose <- 1i

	// 데이터를 읽고 채널을 비우고 채널을 닫음
	<-willClose
	<-willClose
	close(willClose)

	// 빈 채널을 읽어서 Zero Value를 읽음
	read := <-willClose
	fmt.Println(read)
}
