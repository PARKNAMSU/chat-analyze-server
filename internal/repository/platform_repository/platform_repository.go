package platform_repository

import (
	"chat-analyze.com/chat-analyze-server/internal/data_struct/model/platform_model"
	"chat-analyze.com/chat-analyze-server/internal/infra/infra_database"
	"github.com/jmoiron/sqlx"
)

type PlatformRepository struct {
	slaveDB  *sqlx.DB
	masterDB *sqlx.DB
}

var (
	repository *PlatformRepository
)

func GetRepository() *PlatformRepository {
	if repository == nil {
		repository = &PlatformRepository{
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

func (r *PlatformRepository) GetPlatformByDomain(domain string) platform_model.PartnerPlatform {
	platform := platform_model.PartnerPlatform{}

	// todo: GetPlatform logic 구현

	return platform
}
