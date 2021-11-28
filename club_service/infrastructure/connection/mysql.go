package connection

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MysqlConnection struct {
}

func (c MysqlConnection) Get() (*sqlx.DB, error) {
	dbUrl := viper.GetString("db.url")

	db, err := sqlx.Connect("mysql", dbUrl)

	if err != nil {
		log.WithError(err).Error("Can not connect to db.")
		return nil, err
	}

	return db, nil
}
