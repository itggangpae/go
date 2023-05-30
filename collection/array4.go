package main

import (
	"fmt"
)

func main() {
	ar := [3]string{"강진 축구", "사이버 가수 아담", "프리스톤테일"}
	br := ar
	fmt.Println(ar[0])
	fmt.Println(br[0])
	ar[0] = "사이버컵"
	fmt.Println(ar[0])
	fmt.Println(br[0])
}
