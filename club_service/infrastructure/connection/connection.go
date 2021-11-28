package connection

import (
	"github.com/jmoiron/sqlx"
)

type Connection interface {
	Get() (*sqlx.DB, error)
}

var isTest = false

var connection Connection

func Instance() Connection {
	if nil != connection {
		return connection
	}

	if isTest {
		connection = SqliteConnection{connectionIdx: 0}
	} else {
		connection = MysqlConnection{}
	}

	return connection
}

func SetTestEnvironment() {
	isTest = true
}
