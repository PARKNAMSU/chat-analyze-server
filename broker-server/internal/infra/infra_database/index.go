package infra_database

import (
	"log"

	"chat-analyze.com/chat-analyze-server/internal/option"
	"github.com/jmoiron/sqlx"
)

type DBEngine = int
type DBType = int

// 특정 DB에 종속되지 않게 connector interface 추상화
type connectorInterface interface {
	connector() (*sqlx.DB, error)
}

var (
	MYSQL    DBEngine = 0
	POSTGRES DBEngine = 1
)

var (
	SlaveDB  DBType = 0
	MasterDB DBType = 1
)

type ConnectOption struct {
	Engine   DBEngine
	Database DBType
}

func DBConnect(connOption ConnectOption) *sqlx.DB {
	if option.DryRun {
		return nil
	}
	var conn connectorInterface
	// engine 에 따라 처리
	switch connOption.Engine {
	case MYSQL:
		conn = &mysqlDB{
			name: connOption.Database,
		}
	case POSTGRES:
		conn = &postgresDB{
			name: connOption.Database,
		}
	default:
		log.Panicln("not support db engine")
	}
	db, err := conn.connector()
	if err != nil {
		log.Panicln(err)
	}
	return db
}
