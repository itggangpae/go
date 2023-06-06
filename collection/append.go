package main

import (
	"fmt"
)

func main() {
	ar := []int{1, 2, 3}
	fmt.Println("ar[0] 의 참조:", &ar[0])
	ar = append(ar, 4)
	fmt.Println(ar)
	fmt.Println("ar[0] 의 참조:", &ar[0])
	ar = append(ar, 4)
	fmt.Println(ar)
	fmt.Println("ar[0] 의 참조:", &ar[0])

	br := []int{5, 6, 7}
	ar = append(ar, br...)
	fmt.Println(ar)

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a), cap(a))
	a = append(a, 6, 7)
	fmt.Println(len(a), cap(a))

	origin := [3]int{1, 2, 3}
	originslice := origin[0:3]

	origin[0] = 200
	fmt.Println("orijin[0]:", origin[0])
	fmt.Println("originslice[0]:", originslice[0])
	originslice = append(originslice, 49)

	origin[0] = 300
	fmt.Println("orijin[0]:", origin[0])
	fmt.Println("originslice[0]:", originslice[0])
}
