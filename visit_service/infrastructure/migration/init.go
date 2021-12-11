package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var transactions = []string{
	`CREATE TABLE IF NOT EXISTS visit (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    start DATETIME NOT NULL,
    club_id INT UNSIGNED NOT NULL,
    order_details TEXT NOT NULL,
    comment VARCHAR(255) NOT NULL DEFAULT '',
    card_id INT UNSIGNED NOT NULL,
    client_name VARCHAR(255) NOT NULL
);`,
	`ALTER TABLE visit ADD INDEX (club_id);`,
}

func initMigration(db *sqlx.DB) {
	for i := 0; i < len(transactions); i++ {
		log.WithField("migration", transactions[i]).Info("Run migration")
		result := db.MustExec(transactions[i])
		_, err := result.LastInsertId()
		if err != nil {
			log.WithError(err).WithField("migration", transactions[i]).Fatal("Can not execute migration")
		}
	}
}
