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

type dbEngine = string

var (
	mysqlEngine dbEngine = "mysql"
)

type dbConfig struct {
	user             string
	password         string
	host             string
	database         string
	charset          string
	maxAllowedPacket uint
}

type CustomDB struct {
	engine        dbEngine
	conn          *sqlx.DB
	config        dbConfig
	tx            *sqlx.Tx
	isTransaction bool
}

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
