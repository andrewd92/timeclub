package connection

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mlhoyt/ramsql/driver"
	log "github.com/sirupsen/logrus"
)

type RamSqlConnection struct {
}

func (c RamSqlConnection) Get() (*sqlx.DB, error) {
	db, err := sqlx.Connect("ramsql", "Tests")

	if err != nil {
		log.WithError(err).Error("Can not connect to db.")
		return nil, err
	}

	return db, nil
}
