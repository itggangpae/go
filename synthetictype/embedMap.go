package main

import "fmt"

func main() {
	terrestrialPlanet := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022e+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.8676e+24,
			"orbitalPeriod": 224.70069,
		},
		"Earth": map[string]float32{
			"meanRadius":    6371.0,
			"mass":          5.97219e+24,
			"orbitalPeriod": 365.25641,
		},
		"Mars": map[string]float32{
			"meanRadius":    3389.5,
			"mass":          6.4185e+23,
			"orbitalPeriod": 686.9600,
		},
	}

	for key, value := range terrestrialPlanet {
		fmt.Printf("%s\n", key)
		for subKey, subValue := range value {
			fmt.Printf("%s:%f", subKey, subValue)
		}
		fmt.Println()
	}
}
