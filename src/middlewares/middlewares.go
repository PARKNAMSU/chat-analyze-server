package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chat-analyze.com/chat-analyze-server/src/models/common_models"
	"chat-analyze.com/chat-analyze-server/src/tools"
)

// 소켓 연결시점에 공용적으로 처리하는 미들웨어
func CommonMiddleware(next common_models.SocketRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn := tools.GetWebSocket(w, r)
		if conn == nil {
			w.Write([]byte("Connection failed"))
		}

		_, dataStr, err := conn.ReadMessage()

		if err != nil {
			w.Write([]byte("Connection failed"))
		}

		var clientData map[string]int
		err = json.Unmarshal(dataStr, &clientData)
		userId, chatId, err := tools.AttendRoom(r.Header)

		if err != nil || chatId == 0 || userId == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Not Exist Room"))
		}

		tools.PrintInfoLog("AddGroupMiddleware", fmt.Sprintf("Client connected to group: %d\n", chatId))

		next(w, r, &common_models.GetConnectData{
			Conn:   conn,
			UserId: userId,
			ChatId: chatId,
		})
	}
}
