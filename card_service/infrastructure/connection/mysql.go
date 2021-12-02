package connection

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MysqlConnection struct {
}

func (c MysqlConnection) Get() (*sqlx.DB, error) {
	dbUrl := buildUrl()

	db, err := sqlx.Connect("mysql", dbUrl)

	if err != nil {
		log.WithError(err).Error("Can not connect to db.")
		return nil, err
	}

	return db, nil
}

func buildUrl() string {
	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s",
		viper.GetString("db.user"),
		viper.GetString("db.pass"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)
}
