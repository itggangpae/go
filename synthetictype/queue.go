package main

import (
	"container/list"
	"fmt"
)

type Queue struct { //Queue 구조체 정의
	v *list.List
}

func (q *Queue) Push(val interface{}) { //요소 추가
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} { //요소을 반환하면서 삭제
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue() //새로운 큐 생성

	for i := 1; i < 5; i++ { //요소 입력
		queue.Push(i)
	}
	v := queue.Pop()
	for v != nil { //요소 출력
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
