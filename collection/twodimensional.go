package main

import "fmt"

func main() {
	a := [2][5]int{ //이차원 배열 선언
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9}, //여러 줄에 걸쳐 초기화할 때는 마지막에 쉼표를 추가
	}
	for _, arr := range a { //arr값은 순서대로 a[0]의 배열 a[1]의 배열
		for _, v := range arr { //v값은 순서대로 a[0]과 a[1] 배열의 각 원소
			fmt.Print(v, " ") //v값 출력
		}
		fmt.Println()
	}
}
