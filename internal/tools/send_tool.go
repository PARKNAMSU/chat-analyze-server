package tools

import (
	"chat-analyze.com/chat-analyze-server/internal/data_struct/response/common_response"
	"github.com/gorilla/websocket"
)

func SendError(conn *websocket.Conn, message string, status int) {
	conn.WriteJSON(common_response.ResponseDefault{
		Message: &message,
		Status:  status,
	})
}

func SendCheck(conn *websocket.Conn) {
	message := "Alive"
	conn.WriteJSON(common_response.ResponseDefault{
		Message: &message,
		Status:  200,
	})
}
