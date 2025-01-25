package repository_interface

import "chat-platform-api.com/chat-platform-api/src/infra/database"

type RepositoryImpl interface {
	InitRepository()
	GetMasterDB() *database.CustomDB
	GetSlaveDB() *database.CustomDB
	Commit()
	Rollback()
	Connect()
	Close()
	Transaction()
}

type UserRepositoryImpl interface {
	RepositoryImpl
	SetRefreshToken(userId int, token string, deviceId string, ipAddr string) error
	CreateUser(ipAddr string) (int, error)
}

type PlatformRepositoryImpl interface {
	RepositoryImpl
}
