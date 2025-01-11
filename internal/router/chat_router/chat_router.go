package chat_router

import "chat-analyze.com/chat-analyze-server/internal/models/common_models"

func ChatRouter(connData *common_models.GetConnectData, router string) {
	switch router {
	case "SendText":
	case "sendFile":
	case "getAll":
	case "getOne":
	}
}
