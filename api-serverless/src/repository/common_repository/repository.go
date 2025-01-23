package common_repository

import "chat-platform-api.com/chat-platform-api/src/infra/database"

type Repository struct {
	masterDB *database.CustomDB
	slaveDB  *database.CustomDB
}

func (r *Repository) InitRepository() {
	r = &Repository{
		masterDB: database.GetMysqlMaster(true),
		slaveDB:  database.GetMysqlSlave(),
	}
}

func (r *Repository) GetMasterDB() *database.CustomDB {
	return r.masterDB
}

func (r *Repository) GetSlaveDB() *database.CustomDB {
	return r.slaveDB
}

func (r *Repository) Commit() {
	r.masterDB.Commit()
}

func (r *Repository) Rollback() {
	r.masterDB.Rollback()
}

func (r *Repository) Connect() {
	r.masterDB.Connect()
	r.slaveDB.Connect()
}

func (r *Repository) Close() {
	r.masterDB.Close()
	r.slaveDB.Close()
}

func (r *Repository) Transaction() {
	r.masterDB.Transaction()
}
