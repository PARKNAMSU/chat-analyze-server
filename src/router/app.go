package router

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	websocketCreator = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 모든 요청 허용
			return true
		},
	}
)

func App() {
	MessageRouter()
}
