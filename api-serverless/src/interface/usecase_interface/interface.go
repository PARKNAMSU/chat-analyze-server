package usecase_interface

import (
	"chat-platform-api.com/chat-platform-api/src/interface/repository_interface"
	"chat-platform-api.com/chat-platform-api/src/type/model/user_model"
	"chat-platform-api.com/chat-platform-api/src/type/response/user_response"
)

type UseCaseImpl interface {
	ErrorCheck(r repository_interface.RepositoryImpl, err error)
}

type AuthUseCaseImpl interface {
	GenerateToken(user user_model.UserData, deviceId string) (*user_response.UserTokenResponse, error)
}
