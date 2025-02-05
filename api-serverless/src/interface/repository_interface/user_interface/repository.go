package user_interface

import "chat-platform-api.com/chat-platform-api/src/interface/repository_interface/default_interface"

type UserRepositoryImpl interface {
	default_interface.RepositoryImpl
	SetRefreshToken(userId int, token string, deviceId string, ipAddr string) error
	CreateUser(ipAddr string) (int, error)
}
