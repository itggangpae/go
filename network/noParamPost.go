package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func main() {
	// 간단한 http.Post 예제
	reqBody := bytes.NewBufferString("POST 방식")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
