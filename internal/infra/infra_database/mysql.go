package infra_database

import (
	"errors"
	"time"

	"chat-analyze.com/chat-analyze-server/internal/option"
	"github.com/jmoiron/sqlx"
)

type mysqlDB struct {
	name DBType
}

var (
	mysqlSlave  *sqlx.DB
	mysqlMaster *sqlx.DB

	postgresSlave  *sqlx.DB
	postgresMaster *sqlx.DB
)

func (d *mysqlDB) connector() (*sqlx.DB, error) {
	switch d.name {
	case SlaveDB:
		{
			if mysqlSlave != nil {
				return mysqlSlave, nil
			}
			db, err := mysqlSlaveConnector()
			mysqlSlave = db
			return db, err
		}
	case MasterDB:
		{
			if mysqlMaster != nil {
				return mysqlMaster, nil
			}
			db, err := mysqlMasterConnector()
			mysqlMaster = db
			return db, err
		}
	}
	return nil, errors.New("not supported db")
}

func mysqlSlaveConnector() (*sqlx.DB, error) {
	option := option.MysqlSlaveOption()
	db, dbErr := sqlx.Connect(option.Engine, option.User+":"+option.Password+"@tcp("+option.Host+")/"+option.Database+"?charset=utf8mb4&parseTime=True&maxAllowedPacket=0")
	if dbErr != nil {
		return nil, dbErr
	}
	db.SetConnMaxLifetime(time.Minute)
	db.SetConnMaxIdleTime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db, nil
}

func mysqlMasterConnector() (*sqlx.DB, error) {
	option := option.MysqlMasterOption()
	db, dbErr := sqlx.Connect(option.Engine, option.User+":"+option.Password+"@tcp("+option.Host+")/"+option.Database+"?charset=utf8mb4&parseTime=True&maxAllowedPacket=0")
	if dbErr != nil {
		return nil, dbErr
	}
	db.SetConnMaxLifetime(time.Minute)
	db.SetConnMaxIdleTime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db, nil
}
