package chat_usecase

var (
	useCase *ChatUseCase
)

func GetUseCase() *ChatUseCase {
	if useCase == nil {
		useCase = &ChatUseCase{}
	}
	return useCase
}

type ChatUseCase struct{}

// 채팅방 생성 useCase
func (c *ChatUseCase) CreateChat() {}

// 메세지 입력 userCase
func (c *ChatUseCase) SendMessage() {}
