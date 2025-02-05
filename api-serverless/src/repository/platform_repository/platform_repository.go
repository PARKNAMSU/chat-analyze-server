package platform_repository

import (
	"chat-platform-api.com/chat-platform-api/src/repository/common_repository"
)

var (
	repository *PlatformRepository
)

type PlatformRepository struct {
	common_repository.Repository
}

func GetUserRepository() *PlatformRepository {
	if repository == nil {
		repository.InitRepository()
	}
	return repository
}
