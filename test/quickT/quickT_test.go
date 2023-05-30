package quickT

import (
	"testing"
	"testing/quick"
)

var N = 1000000

func TestWithItself(t *testing.T) {
	condition := func(a, b Point2D) bool {
		return Add(a, b) == Add(b, a)
	}
	//첫번째 함수의 시그니처에 따라 랜덤한 값을 자동으로 생성
	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

// 실행에 실패하는 함수
func TestFail(t *testing.T) {
	condition := func(a, b, c Point2D) bool {
		return Add(Add(a, b), c) == Add(a, b)
	}

	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
