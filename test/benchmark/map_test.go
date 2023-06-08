package main

import (
	"testing"
)

func test(m map[int]int) {
	for i := 0; i < 100000; i++ {
		m[i] = i
	}
}

// 초기 용량을 설정하지 않고 Map 사용
func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		test(m)
	}
}

// 초기 용량을 설정한 후 Map 사용
func BenchmarkCapacityMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, 100000)
		test(m)
	}
}
