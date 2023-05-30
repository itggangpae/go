package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var PORT = ":1234"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!\n")
	fmt.Fprintf(w, "Please use /ws for WebSocket!")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from:", r.Host)

	//HTTP 요청에서 웹 소켓 연결로 업그레이드
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrader.Upgrade:", err)
		return
	}
	defer ws.Close()

	//연결에 성공하면 클라이언트와 동신 시작
	for {
		//메시지를 받는데 성공하면
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("From", r.Host, "read", err)
			break
		}
		//받은 메시지 출력
		log.Print("Received: ", string(message))
		//받은 메시지를 전송
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("WriteMessage:", err)
			break
		}
	}
}

func main() {
	//HTTP 요청 멀티플렉서 인스턴스 생성
	mux := http.NewServeMux()
	//상세 정보 설정
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	//웹 소켓으로 사용할 EndPoint 설정 - ws 대신에 다른 문자열도 가능
	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/ws", http.HandlerFunc(wsHandler))

	log.Println("Listening to TCP Port", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
