package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var transactions = []string{
	`CREATE TABLE IF NOT EXISTS visit (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    start DATETIME NOT NULL,
    club_id INT UNSIGNED NOT NULL,
    order_details TEXT NOT NULL,
    comment VARCHAR(255) NOT NULL DEFAULT '',
    card_id INT UNSIGNED NOT NULL,
    client_name VARCHAR(255) NOT NULL,
	INDEX (club_id)
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
