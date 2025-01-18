package chat_router

import (
	"net/http"

	"chat-analyze.com/chat-analyze-server/internal/controller/chat_controller"
	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
	"chat-analyze.com/chat-analyze-server/internal/middleware/common_middleware"
	"chat-analyze.com/chat-analyze-server/internal/middleware/index_middleware"
)

var (
	chatRouter *http.ServeMux
	controller = chat_controller.GetController()
)

func WSChatRouter(connData *common_model.GetConnectData, router string) {
	switch router {
	case "SendText":
		controller.SendMessage()
	case "sendFile":
	case "getAll":
	case "getOne":
	}
}

func APIChatRouter() *http.ServeMux {
	if chatRouter == nil {
		chatRouter := http.NewServeMux()
		chatRouter.HandleFunc("/createChat", index_middleware.MiddlewareChaining(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("created"))
		}, common_middleware.PlatformValidation))
	}
	return chatRouter
}
