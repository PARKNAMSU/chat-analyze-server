package platform_usecase

import (
	"errors"

	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/platform_model"
	"chat-analyze.com/chat-analyze-server/internal/option"
	"chat-analyze.com/chat-analyze-server/internal/repository/platform_repository"
	"chat-analyze.com/chat-analyze-server/internal/tools"
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

func (p *PlatformUseCase) APIKeyValidation(domain string, APIKey string) (*platform_model.PartnerPlatform, error) {
	platform := p.platformRepository.GetPlatformByDomain(domain)
	keyBytes, err := tools.Decrypt(platform.ApiKey(), option.API_ENCRIPTION_KEY)

	if err != nil {
		return nil, errors.New("API Key Decryption Error")
	}

	if string(keyBytes) != APIKey {
		return nil, errors.New(option.INVALID_AUTHORIZATION)
	}
	return &platform, nil
}
