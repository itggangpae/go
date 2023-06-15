package main

import (
	"fmt"
	"sort"
	"strings"
)

// 사용자 정의 정렬을 위한 메서드 구현
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return strings.ToUpper(s[i]) < strings.ToUpper(s[j])
}

func main() {
	//배열 생성
	sInts := []int{1, 0, 2, -3, 4, -20}
	sStrings := []string{"ggangpae1", "adam", "itstudy", "ITggangpae"}

	//오름차순 정렬
	fmt.Println("sInts original:", sInts)
	sort.Ints(sInts)
	fmt.Println("sInts:", sInts)

	//내림차순 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(sInts)))
	fmt.Println("Reverse:", sInts)

	fmt.Println("sStrings original:", sStrings)
	sort.Strings(sStrings)
	fmt.Println("sStrings:", sStrings)
	sort.Sort(sort.Reverse(sort.StringSlice(sStrings)))
	fmt.Println("Reverse:", sStrings)

	//대소문자 구분없이 정렬
	sort.Sort(ByLength(sStrings))
	fmt.Println("sStrings:", sStrings)
}
