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
	req := httptest.NewRequest("GET", "/student", nil) // /student 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	student := new(Student)
	err := json.NewDecoder(res.Body).Decode(student) //결과 변환
	assert.Nil(err)                                  //결과 확인
	assert.Equal("adam", student.Name)
	assert.Equal(53, student.Age)
	assert.Equal(99, student.Score)
}
