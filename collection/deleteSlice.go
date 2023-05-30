package main

import (
	"fmt"
)

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	i := 1
	//삭제 가능한 인덱스인지 확인
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	//삭제 위치 전까지 와 삭제 위치 다음부터의 슬라이스 결합
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("After 1st deletion:", aSlice)
}
