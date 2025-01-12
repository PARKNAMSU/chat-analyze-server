package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"chat-analyze.com/chat-analyze-server/internal/data_struct/dto/common_dto"
	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
	"chat-analyze.com/chat-analyze-server/internal/middlewares"
	"chat-analyze.com/chat-analyze-server/internal/options"
	"chat-analyze.com/chat-analyze-server/internal/router/chat_router"
	"chat-analyze.com/chat-analyze-server/internal/tools"
	"github.com/joho/godotenv"
)

var (
	_    = godotenv.Load()
	port = os.Getenv("PORT")
)

func App() {
	appHandler := func(w http.ResponseWriter, r *http.Request, connData *common_model.GetConnectData) {
		conn := connData.Conn
		defer conn.Close()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				break
			}

			var clientData common_dto.DefaultRequest

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

	http.HandleFunc("/ws", middlewares.SocketMiddleware(appHandler))

	err := http.ListenAndServe(fmt.Sprintf(":"+port), nil)

	if err != nil {
		tools.PanicError("App", fmt.Sprintf("Server start error: %s", err.Error()))
		return
	}

	tools.PrintInfoLog("App", "Server started at "+port)
}
