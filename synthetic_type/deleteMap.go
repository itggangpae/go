package main

import (
	"fmt"
)

func main() {
	a := map[string]string{"name": "adam", "age": "51"}
	delete(a, "age")
	for key, value := range a {
		fmt.Printf("%s:%s\n", key, value)
	}
}
