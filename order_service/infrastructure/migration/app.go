package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var transactions = []string{`
CREATE TABLE IF NOT EXISTS orders (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	status VARCHAR(255) NOT NULL DEFAULT 'OPEN',
	start DATETIME NOT NULL,
	end DATETIME
);`,
	`CREATE TABLE IF NOT EXISTS order_visits (
      id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
      order_id INT UNSIGNED NOT NULL,
      visit_id INT UNSIGNED NOT NULL,
      index (order_id)
);`,
}

func initMigration(db *sqlx.DB) {
	firstTransaction := 0
	for i := firstTransaction; i < len(transactions); i++ {
		log.WithField("migration", transactions[i]).Info("Run migration")
		result := db.MustExec(transactions[i])
		_, err := result.LastInsertId()
		if err != nil {
			log.WithError(err).WithField("migration", transactions[i]).Fatal("Can not execute migration")
		}
	}
}
