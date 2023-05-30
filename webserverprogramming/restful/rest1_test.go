package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil) ///students 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list) // 결과 변환
	assert.Nil(err)                                // 결과 확인
	assert.Equal(3, len(list))
	assert.Equal("adam", list[0].Name)
	assert.Equal("jessica", list[1].Name)
	assert.Equal("hunt", list[2].Name)
}
