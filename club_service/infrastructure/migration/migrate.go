package migration

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() {
	dbUrl := viper.GetString("db.url")

	db, err := sqlx.Connect("mysql", dbUrl)

	if err != nil {
		log.WithError(err).Fatal("Can not connect to db")
	}

	initMigration(db)
}
