package main

import (
	"fmt"
)

func main() {
	a := map[string]string{"name": "adam", "age": "51"}
	fmt.Println(a["name"])
	_, ok := a["name"]
	fmt.Println(ok)

	fmt.Println(a["address"])
	_, ok = a["address"]
	fmt.Println(ok)
}
