package main

import "fmt"

func main() {
	//2차원 슬라이스
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	//2차원 슬라이스 순회
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)
}
