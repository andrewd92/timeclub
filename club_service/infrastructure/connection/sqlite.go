package connection

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

const poolSize = 20

var connectionIdx = 0

var connections []*sqlx.DB

type SqliteConnection struct {
	connectionIdx int
}

func (c SqliteConnection) Get() (*sqlx.DB, error) {
	if connections == nil {
		connections = make([]*sqlx.DB, poolSize, poolSize)
		for i := 0; i < poolSize; i++ {
			db, err := sqlx.Connect("sqlite3", "file:test?mode=memory&cache=shared")

			if err != nil {
				log.WithError(err).Error("Can not connect to db.")
				return nil, err
			}

			connections[i] = db
		}
	}

	con := connections[connectionIdx]
	connectionIdx += 1

	return con, nil
}
