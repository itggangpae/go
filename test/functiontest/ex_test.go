package main

import (
	"testing"
)

func TestSquare1(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) 81이어야 하는데 %d 값을 리턴함\n", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("square(3) 9이어야 하는데 %d 값을 리턴함\n", rst)
	}
}
