package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t) //테스트 객체 생성
	//함수 호출 결과 확인
	assert.Equal(81, square(9), "square(9) 의 결과는 81") //테스트 함수 호출
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) 의 결과는 9")
}
