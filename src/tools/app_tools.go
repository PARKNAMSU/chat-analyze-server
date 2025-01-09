package tools

import (
	"log"
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

func GetWebSocket(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	conn, err := websocketCreator.Upgrade(w, r, nil)

	if err != nil {
		PrintErrorLog("getWebSocket", err.Error())
		return nil
	}
	log.Println("Client connected")
	return conn
}
