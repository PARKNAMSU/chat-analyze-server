package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"chat-analyze.com/chat-analyze-server/src/middlewares"
	"chat-analyze.com/chat-analyze-server/src/models/common_models"
	"chat-analyze.com/chat-analyze-server/src/models/messaging_models"
	"chat-analyze.com/chat-analyze-server/src/options"
	"chat-analyze.com/chat-analyze-server/src/router/chat_router"
	"chat-analyze.com/chat-analyze-server/src/tools"
	"github.com/gorilla/websocket"
)

var (
	websocketCreator = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 모든 요청 허용
			return true
		},
	}

	port = ":8080"
)

func App() {
	appHandler := func(w http.ResponseWriter, r *http.Request, connData *common_models.GetConnectData) {
		conn := connData.Conn
		defer conn.Close()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				break
			}

			var clientData messaging_models.RequestDefault

			err = json.Unmarshal(message, &clientData)

			if err != nil {
				tools.SendError(conn, options.INVALID_ROUTER, options.StatusBadRequest)
				continue
			}

			routers := strings.Split(clientData.Router, "/")

			if len(routers) < 3 {
				tools.SendError(conn, options.INVALID_ROUTER, options.StatusBadRequest)
				continue
			}

			switch routers[1] {
			case "check":
				tools.SendCheck(conn)
			case "chat":
				chat_router.ChatRouter(connData, routers[2])
			default:
				tools.SendError(conn, options.INVALID_ROUTER, options.StatusBadRequest)
			}

		}
	}

	http.HandleFunc("/ws", middlewares.CommonMiddleware(appHandler))

	err := http.ListenAndServe(port, nil)

	if err != nil {
		tools.PanicError("App", fmt.Sprintf("Server start error: %s", err.Error()))
		return
	}

	tools.PrintInfoLog("App", "Server started at "+port)
}
