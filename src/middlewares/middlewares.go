package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chat-analyze.com/chat-analyze-server/src/cache"
	"chat-analyze.com/chat-analyze-server/src/models"
	"chat-analyze.com/chat-analyze-server/src/tools"
	"github.com/gorilla/websocket"
)

func AddGroupMiddleware(next models.SocketRouter) http.HandlerFunc {
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
		roomId, isExist := clientData["roomId"]

		if err != nil || !isExist || roomId == 0 {
			w.Write([]byte("Not Exist RoomId"))
		}

		if cache.USER_GROUP[roomId] == nil {
			cache.USER_GROUP[roomId] = make(map[*websocket.Conn]bool)
		}

		cache.USER_GROUP[roomId][conn] = true

		tools.PrintInfoLog("AddGroupMiddleware", fmt.Sprintf("Client connected to group: %d\n", clientData["roomId"]))

		next(w, r, conn)
	}
}
