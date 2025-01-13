package infra_database

import (
	"errors"
	"fmt"
	"time"

	"chat-analyze.com/chat-analyze-server/internal/option"
	"github.com/jmoiron/sqlx"
)

type postgresDB struct {
	name DBName
}

func (d *postgresDB) connector() (*sqlx.DB, error) {
	switch d.name {
	case SlaveDB:
		{
			if postgresSlave != nil {
				return postgresSlave, nil
			}
			db, err := postgresSlaveConnector()
			postgresSlave = db
			return db, err
		}
	case MasterDB:
		{
			if postgresMaster != nil {
				return postgresMaster, nil
			}
			db, err := postgresMasterConnector()
			postgresMaster = db
			return db, err
		}
	}
	return nil, errors.New("not supported db")
}

func postgresSlaveConnector() (*sqlx.DB, error) {
	option := option.PostgresSlaveOption()
	db, dbErr := sqlx.Connect(option.Engine, fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", option.User, option.Password, option.Host, option.Database))
	if dbErr != nil {
		return nil, dbErr
	}
	db.SetConnMaxLifetime(time.Minute)
	db.SetConnMaxIdleTime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db, dbErr
}

func postgresMasterConnector() (*sqlx.DB, error) {
	option := option.PostgresMasterOption()
	db, dbErr := sqlx.Connect(option.Engine, fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", option.User, option.Password, option.Host, option.Database))
	if dbErr != nil {
		return nil, dbErr
	}
	db.SetConnMaxLifetime(time.Minute)
	db.SetConnMaxIdleTime(time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db, dbErr
}
