package database

import (
	"fmt"
	"os"
)

var (
	mysqlMasterDB *CustomDB
	mysqlSlaveDB  *CustomDB
)

func getMysqlConnect(config dbConfig) string {
	return fmt.Sprintf(
		`%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&maxAllowedPacket=%d`,
		config.user,
		config.password,
		config.host,
		config.database,
		config.charset,
		config.maxAllowedPacket,
	)
}

func InitMysqlMaster(isTransaction bool) {
	if mysqlMasterDB != nil {
		return
	}
	mysqlMasterDB = &CustomDB{
		engine: mysqlEngine,
		config: dbConfig{
			user:             os.Getenv("MYSQL_USER_MASTER"),
			password:         os.Getenv("MYSQL_PASSWORD_MASTER"),
			host:             os.Getenv("MYSQL_HOST_MASTER"),
			database:         os.Getenv("MYSQL_DATABASE_MASTER"),
			charset:          "utf8mb4",
			maxAllowedPacket: 0,
		},
		isTransaction: isTransaction,
	}
	mysqlMasterDB.Connect()
}

func InitMysqlSlave(isTransaction bool) {
	if mysqlSlaveDB != nil {
		return
	}
	mysqlSlaveDB = &CustomDB{
		engine: mysqlEngine,
		config: dbConfig{
			user:             os.Getenv("MYSQL_USER_MASTER"),
			password:         os.Getenv("MYSQL_PASSWORD_MASTER"),
			host:             os.Getenv("MYSQL_HOST_MASTER"),
			database:         os.Getenv("MYSQL_DATABASE_MASTER"),
			charset:          "utf8mb4",
			maxAllowedPacket: 0,
		},
		isTransaction: isTransaction,
	}
	mysqlSlaveDB.Connect()
}

func GetMysqlMaster() *CustomDB {
	return mysqlMasterDB
}

func GetMysqlSlave() *CustomDB {
	return mysqlSlaveDB
}
