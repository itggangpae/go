package main

import "fmt"

func main() {
	haitai := []string{"이대진", "최향남", "이종범", "선동열", "김상진"}
	kia := []string{"윤석민", "양현종", "이범호", "김주찬"}

	kbo := map[string][]string{
		"haital": haitai,
		"kia":    kia,
	}

	for key, value := range kbo {
		fmt.Printf("%s\n", key)
		for _, player := range value {
			fmt.Printf("%s\t", player)
		}
		fmt.Println()
	}
}
