package chat_repository

import (
	"chat-analyze.com/chat-analyze-server/internal/infra/infra_database"
	"github.com/jmoiron/sqlx"
)

type ChatRepository struct {
	slaveDB  *sqlx.DB
	masterDB *sqlx.DB
}

var (
	repository *ChatRepository
)

func GetRepository() *ChatRepository {
	if repository == nil {
		repository = &ChatRepository{
			slaveDB: infra_database.
				DBConnect(
					infra_database.
						ConnectOption{
						Engine:   infra_database.MYSQL,
						Database: infra_database.SlaveDB,
					},
				),
			masterDB: infra_database.
				DBConnect(
					infra_database.
						ConnectOption{
						Engine:   infra_database.MYSQL,
						Database: infra_database.MasterDB,
					},
				),
		}
	}
	return repository
}

func (r *ChatRepository) SendMessage() {
	// do something
}
