package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()     // 연결 리스트 생성
	l.PushBack("강진 축구") // 연결 리스트에 데이터 추가
	l.PushBack("아담")
	l.PushBack("사이버 컵")
	l.PushBack("프리스톤테일")

	fmt.Println("Front ", l.Front().Value) // 연결 리스트의 맨 앞 데이터를 가져옴
	fmt.Println("Back ", l.Back().Value)   // 연결 리스트의 맨 뒤 데이터를 가져옴

	for e := l.Front(); e != nil; e = e.Next() { // 연결 리스트의 맨 앞부터 끝까지 순회
		fmt.Println(e.Value)
	}
}
