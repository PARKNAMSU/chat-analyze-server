package chat_router

import (
	"chat-analyze.com/chat-analyze-server/internal/controller/chat_controller"
	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/common_model"
)

var (
	controller = chat_controller.GetController()
)

func ChatRouter(connData *common_model.GetConnectData, router string) {
	switch router {
	case "SendText":
		controller.SendMessage()
	case "sendFile":
	case "getAll":
	case "getOne":
	}
}
