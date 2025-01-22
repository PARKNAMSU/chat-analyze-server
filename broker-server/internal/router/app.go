package router

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"chat-analyze.com/chat-analyze-server/internal/cache"
	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/chat_model"
	"chat-analyze.com/chat-analyze-server/internal/infra"
	"chat-analyze.com/chat-analyze-server/internal/middleware/common_middleware"
	"chat-analyze.com/chat-analyze-server/internal/middleware/index_middleware"
	"chat-analyze.com/chat-analyze-server/internal/option"
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

	connData, err := tools.GetConnData(w, r)

	if err != nil {
		tools.SendErrorResponse(
			w,
			option.INTERNAL_SERVER_ERROR,
			option.StatusInternalServerError,
		)
		return
	}

	conn := connData.Conn
	cache.SetChatCache(connData)

	defer conn.Close()

	msgChan := make(chan chat_model.MessageData)

	consumer := infra.KafkaSubscribeTopic(strconv.Itoa(connData.ChatId))

	if consumer == nil {
		tools.SendErrorResponse(
			w,
			option.INTERNAL_SERVER_ERROR,
			option.StatusInternalServerError,
		)
		return
	}

	go infra.KafkaPolling(consumer, msgChan)

	for messages := range msgChan {
		tools.WSSendMessage(connData, messages)
	}
}

func App() {

	mux := http.NewServeMux()

	mux.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("alive"))
	})

	mux.HandleFunc("/ws", index_middleware.MiddlewareChaining(socketHandler))

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
