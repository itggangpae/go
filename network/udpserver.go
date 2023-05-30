package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	//포트 번호 설정
	PORT := ":1234"

	//앤드포인트 설정
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	//UDP 서버 생성
	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	//데이터를 읽기 위한 버퍼 생성
	buffer := make([]byte, 1024)
	//랜덤한 숫자 출력을 위한 시드 설정
	rand.Seed(time.Now().Unix())

	for {
		//UDPClient로부터 전송된 데이터를 읽어서 출력
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		//Client 가 STOP 이라는 메세지가 오면 서버 종료
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		//클라이언트에게 받아온 데이터 전송
		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
