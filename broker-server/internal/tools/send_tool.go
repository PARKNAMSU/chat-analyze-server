package tools

import (
	"encoding/json"
	"net/http"

	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
	"chat-analyze.com/chat-analyze-server/internal/data_struct/response/common_response"
	"github.com/gorilla/websocket"
)

func WSSendError(conn *websocket.Conn, message string, status int) {
	conn.WriteJSON(common_response.ResponseDefault{
		Message: &message,
		Status:  status,
	})
}

func SendErrorResponse(w http.ResponseWriter, message string, status int) {
	res := common_response.ResponseDefault{
		Message: &message,
		Status:  status,
	}

	w.WriteHeader(http.StatusUnauthorized)
	data, _ := json.Marshal(res)

	w.Write(data)
}

func WSSendCheck(conn *websocket.Conn) {
	message := "Alive"
	conn.WriteJSON(common_response.ResponseDefault{
		Message: &message,
		Status:  200,
	})
}

func WSSendMessage[T any](conn *common_model.GetConnectData, responseData T) {
	conn.Conn.WriteJSON(map[string]any{
		"userId": conn.UserId,
		"chatId": conn.ChatId,
		"data":   responseData,
	})
}
