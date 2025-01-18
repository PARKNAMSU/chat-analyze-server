package chat_usecase

import "chat-analyze.com/chat-analyze-server/internal/repository/chat_repository"

var (
	useCase *ChatUseCase
)

func GetUseCase() *ChatUseCase {
	if useCase == nil {
		useCase = &ChatUseCase{
			chatRepository: chat_repository.GetRepository(),
		}
	}
	return useCase
}

type ChatUseCase struct {
	chatRepository *chat_repository.ChatRepository
}

// 채팅방 생성 useCase
func (c *ChatUseCase) CreateChat() {}

// 메세지 입력 userCase
func (c *ChatUseCase) SendMessage() {
	c.chatRepository.SendMessage()
}
