package chat_controller

import (
	"chat-analyze.com/chat-analyze-server/internal/usecase/chat_usecase"
)

var (
	controller *ChatController
)

func GetController() *ChatController {
	if controller == nil {
		controller = &ChatController{}
	}
	return controller
}

type ChatController struct {
	chatUseCase *chat_usecase.ChatUseCase
}

func (c *ChatController) SendMessage() {
	c.chatUseCase.SendMessage()
}
