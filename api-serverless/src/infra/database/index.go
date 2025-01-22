package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"chat-platform-api.com/chat-platform-api/src/variable/common_variable"
	"github.com/jmoiron/sqlx"
)

/*
db engine 타입.

description:

	확장성 있는 db 연결을 위해 해당 타입으로 DBMS 종류 정의
*/
type dbEngine = string

/* db connect 시 설정 데이터 타입 */
type dbConfig struct {
	user             string
	password         string
	host             string
	database         string
	charset          string
	maxAllowedPacket uint
}

/*
외부에서 사용할 DB Connect.

description:

	내부 인자 직접 변경 및 사용할 수 없게 private 하게 설정
*/
type CustomDB struct {
	engine        dbEngine // DBMS 타입
	conn          *sqlx.DB // connection 객체
	config        dbConfig // db 설정
	tx            *sqlx.Tx // transaction 객체
	isTransaction bool     // transaction 사용 여부
}

var (
	mysqlEngine dbEngine = "mysql"
)

func (c *CustomDB) Connect() {
	if c.conn != nil {
		return
	}
	connector := ""

	switch c.engine {
	case mysqlEngine:
		connector = getMysqlConnect(c.config)
	default:
	}

	db, err := sqlx.Connect(
		string(c.engine),
		connector,
	)

	if err != nil {
		log.Println(err.Error())
		return
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetConnMaxIdleTime(time.Minute)

	maxIdleConn := 5
	maxOpenConn := 5
	switch common_variable.ENVIRONMENT {
	case "production":
		maxIdleConn = 15
		maxOpenConn = 15
	case "staging":
		maxIdleConn = 10
		maxOpenConn = 10
	}

	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	c.conn = db
}

func (c *CustomDB) Transaction() {
	if tx, err := c.conn.BeginTxx(context.TODO(), nil); err != nil {
		fmt.Println("[err]:[setTransaction]:[", err.Error(), "]")
		c.tx = tx
	}
}

func (c *CustomDB) Commit() {
	if c.tx == nil {
		return
	}
	if err := c.tx.Commit(); err != nil {
		fmt.Println("[err]:[Commit]:[", err.Error(), "]")
	}
	c.tx = nil
}

func (c *CustomDB) Rollback() {
	if c.tx == nil {
		return
	}
	if err := c.tx.Rollback(); err != nil {
		fmt.Println("[err]:[Rollback]:[", err.Error(), "]")
	}
	c.tx = nil
}

func (c *CustomDB) QueryExecute(query string, queryParams []any) (sql.Result, error) {
	if c.isTransaction {
		return c.tx.Exec(query, queryParams...)
	}
	return c.conn.Exec(query, queryParams...)
}

func (c *CustomDB) QuerySelect(query string, queryParams []any) any {
	var data any

	if c.isTransaction {
		c.tx.Select(&data, query, queryParams...)
	}
	c.conn.Select(&data, query, queryParams...)

	return data
}
