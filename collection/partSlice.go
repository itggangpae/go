package main

import "fmt"

func main() {
	//슬라이스 생성
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	l := len(aSlice)

	//앞에서부터 5개 elements
	fmt.Println(aSlice[0:5])
	//앞에서부터 5개 elements
	fmt.Println(aSlice[:5])

	//마지막 2개 elements
	fmt.Println(aSlice[l-2 : l])
	//마지막 2개 elements
	fmt.Println(aSlice[l-2:])

	//앞에서 5개 elements
	t := aSlice[0:5:10]
	fmt.Println(len(t), cap(t))

	//2,3,4 번째 인덱스
	t = aSlice[2:5:10]
	fmt.Println(len(t), cap(t))

	//0,1,2,3,4 번째 인덱스
	t = aSlice[:5:6]
	fmt.Println(len(t), cap(t))
}
