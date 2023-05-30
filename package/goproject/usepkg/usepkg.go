package main

import (
	"fmt"
	"goproject/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
)

func main() {
	custompkg.PrintCustom()
	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
