package user_repository

import (
	"chat-platform-api.com/chat-platform-api/src/repository/common_repository"
)

var (
	repository *UserRepository
)

type UserRepository struct {
	common_repository.Repository
}

func GetUserRepository() *UserRepository {
	if repository == nil {
		repository.InitRepository()
	}
	return repository
}

func (r *UserRepository) SetRefreshToken(userId int, token string, deviceId string, ipAddr string) error {
	_, err := r.GetMasterDB().NamedQueryExecute(
		"INSERT INTO `user_refresh_token` SET `userId` = :userId , `token` = :token , `deviceId` = :deviceId , ipAddr = :ipAddr "+
			"ON DUPLICATE KEY UPDATE `token` = :token, `deviceId` = :deviceId, `ipAddr` = :ipAddr",
		map[string]any{
			"userId":   userId,
			"token":    token,
			"deviceId": deviceId,
			"ipAddr":   ipAddr,
		},
	)

	if err != nil {
		return err
	}
	return nil
}
