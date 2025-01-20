package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"chat-analyze.com/chat-analyze-server/internal/data_struct/dto/common_dto"
	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
	"chat-analyze.com/chat-analyze-server/internal/middleware/chat_middleware"
	"chat-analyze.com/chat-analyze-server/internal/middleware/common_middleware"
	"chat-analyze.com/chat-analyze-server/internal/middleware/index_middleware"
	"chat-analyze.com/chat-analyze-server/internal/option"
	"chat-analyze.com/chat-analyze-server/internal/router/chat_router"
	"chat-analyze.com/chat-analyze-server/internal/tools"
	"github.com/joho/godotenv"
)

var (
	_    = godotenv.Load()
	port = os.Getenv("PORT")
)

func socketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		tools.SendErrorResponse(w, option.NOT_FOUND, option.StatusNotFound)
	}
	conn := tools.GetWebSocket(w, r)

	userId, isUserExist := r.Context().Value(option.CONTEXT_USER_ID).(int)
	chatId, isChatExist := r.Context().Value(option.CONTEXT_CHAT_ID).(int)

	if !isUserExist || !isChatExist {
		tools.WSSendError(conn, option.INVALID_ROUTER, option.StatusBadRequest)
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
			tools.WSSendError(conn, option.INVALID_ROUTER, option.StatusBadRequest)
			continue
		}

		routers := strings.Split(clientData.Router, "/")

		if len(routers) < 3 {
			tools.WSSendError(conn, option.INVALID_ROUTER, option.StatusBadRequest)
			continue
		}

		switch routers[1] {
		case "check":
			tools.WSSendCheck(conn)
		case "chat":
			chat_router.WSChatRouter(connData, routers[2])
		default:
			tools.WSSendError(conn, option.INVALID_ROUTER, option.StatusBadRequest)
		}
	}
}

func App() {

	mux := http.NewServeMux()

	mux.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("alive"))
	})

	mux.HandleFunc("/ws", index_middleware.MiddlewareChaining(socketHandler, common_middleware.PlatformValidation, chat_middleware.AttendChatMiddleware))

	mux.HandleFunc("/restartTest", func(w http.ResponseWriter, r *http.Request) {
		log.Panicln("server down")
		w.Write([]byte("server down"))
	})

	tools.PrintInfoLog("App", "Broker Server started at "+port)

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
		tools.PanicError("App", fmt.Sprintf("Broker Server start error: %s", err.Error()))
		return
	}
}
