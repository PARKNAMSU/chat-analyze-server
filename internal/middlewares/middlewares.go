package middlewares

import (
	"fmt"
	"net/http"

	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
	"chat-analyze.com/chat-analyze-server/internal/tools"
)

// 소켓 연결시점에 공용적으로 처리하는 미들웨어
func SocketMiddleware(next common_model.SocketRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn := tools.GetWebSocket(w, r)
		if conn == nil {
			w.Write([]byte("Connection failed"))
		}

		userId, chatId, err := tools.AttendRoom(r.Header)

		if err != nil || chatId == 0 || userId == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Not Exist Room"))
		}

		tools.PrintInfoLog("AddGroupMiddleware", fmt.Sprintf("Client connected to group: %d\n", chatId))

		next(w, r, &common_model.GetConnectData{
			Conn:   conn,
			UserId: userId,
			ChatId: chatId,
		})
	}
}
