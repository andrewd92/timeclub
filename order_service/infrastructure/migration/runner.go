package migration

import (
	"github.com/andrewd92/timeclub/order_service/infrastructure/connection"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
)

func Run() {
	db, err := connection.Instance().Get()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.WithError(err).Error("Can not close connection in migrations")
		}
	}(db)

	if err != nil {
		log.WithError(err).Fatal("Can not connect to db")
	}

	configName := os.Getenv("VIPER_CONFIG_NAME")

	if configName == "test" {
		initTestMigration(db)
	} else {
		initMigration(db)
	}
}
