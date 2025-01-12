package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"chat-analyze.com/chat-analyze-server/internal/data_struct/dto/common_dto"
	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
	"chat-analyze.com/chat-analyze-server/internal/middleware/chat_middleware"
	"chat-analyze.com/chat-analyze-server/internal/middleware/common_middleware"
	"chat-analyze.com/chat-analyze-server/internal/middleware/index_middleware"
	"chat-analyze.com/chat-analyze-server/internal/options"
	"chat-analyze.com/chat-analyze-server/internal/router/chat_router"
	"chat-analyze.com/chat-analyze-server/internal/tools"
	"github.com/joho/godotenv"
)

var (
	_    = godotenv.Load()
	port = os.Getenv("PORT")
)

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn := tools.GetWebSocket(w, r)

	userId, isUserExist := r.Context().Value(options.CONTEXT_USER_ID).(int)
	chatId, isChatExist := r.Context().Value(options.CONTEXT_CHAT_ID).(int)

	if !isUserExist || !isChatExist {
		tools.SendError(conn, options.INVALID_ROUTER, options.StatusBadRequest)
		return
	}

	connData := &common_model.GetConnectData{
		Conn:   conn,
		UserId: userId,
		ChatId: chatId,
	}

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

func App() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ws", index_middleware.MiddlewareChaining(socketHandler, chat_middleware.AttendChatMiddleware))
	mux.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("alive"))
	})

	tools.PrintInfoLog("App", "Server started at "+port)

	err := http.ListenAndServe(
		fmt.Sprintf(":"+port),
		http.HandlerFunc( // 서버 접속 시 공용으로 처리하는 미들웨어
			index_middleware.MiddlewareChaining(
				mux.ServeHTTP,
				common_middleware.SetHeader,
				common_middleware.APIKeyValidation,
			),
		),
	)

	if err != nil {
		tools.PanicError("App", fmt.Sprintf("Server start error: %s", err.Error()))
		return
	}
}
