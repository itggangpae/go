package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    //쿼리 인수 가져오기
	name := values.Get("name") //특정 키값이 있는지 확인
	if name == "" {
		name = "noname"
	}
	id, _ := strconv.Atoi(values.Get("id")) //id값을 가져와서 int타입 변환
	fmt.Fprintf(w, "반갑습니다 %s! id:%d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler) //"/bar" 핸들러 등록
	http.ListenAndServe(":3000", nil)
}
