package tools

import (
	"chat-analyze.com/chat-analyze-server/src/models/messaging_models"
	"github.com/gorilla/websocket"
)

func SendError(conn *websocket.Conn, message string, status int) {
	conn.WriteJSON(messaging_models.ResponseDefault{
		Message: &message,
		Status:  status,
	})
}

func SendCheck(conn *websocket.Conn) {
	message := "Alive"
	conn.WriteJSON(messaging_models.ResponseDefault{
		Message: &message,
		Status:  200,
	})
}
