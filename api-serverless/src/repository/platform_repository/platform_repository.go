package platform_repository

import (
	"chat-platform-api.com/chat-platform-api/src/interface/repository_interface"
	"chat-platform-api.com/chat-platform-api/src/repository/common_repository"
)

var (
	repository *PlatformRepository
)

type PlatformRepository struct {
	common_repository.Repository
}

func GetUserRepository() repository_interface.PlatformRepositoryImpl {
	if repository == nil {
		repository.InitRepository()
	}
	return repository
}
