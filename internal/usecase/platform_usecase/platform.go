package platform_usecase

import (
	"chat-analyze.com/chat-analyze-server/internal/repository/platform_repository"
)

type PlatformUseCase struct {
	platformRepository *platform_repository.PlatformRepository
}

var (
	useCase *PlatformUseCase
)

func GetUseCase() *PlatformUseCase {
	if useCase == nil {
		useCase = &PlatformUseCase{
			platformRepository: platform_repository.GetRepository(),
		}
	}
	return useCase
}

func (p *PlatformUseCase) APIKeyValidation() {

}
