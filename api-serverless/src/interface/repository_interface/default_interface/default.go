package default_interface

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
