package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var testTransactions = []string{
	`CREATE TABLE IF NOT EXISTS visit (
    id INTEGER NOT NULL PRIMARY KEY,
    start TEXT NOT NULL,
    club_id INTEGER NOT NULL,
    order_details TEXT NOT NULL,
    comment VARCHAR(255) NOT NULL DEFAULT '',
    card_id INTEGER NOT NULL,
    client_name VARCHAR(255) NOT NULL
);`,
	`INSERT INTO visit(id, start, club_id, order_details, comment, card_id, client_name)
VALUES
   (1, datetime('now', '-1 Hour'), 1, '[]', '', 1, 'Sasha');`,
}

func initTestMigration(db *sqlx.DB) {
	firstTransaction := 0
	for i := firstTransaction; i < len(testTransactions); i++ {
		log.WithField("migration", testTransactions[i]).Info("Run migration")
		result := db.MustExec(testTransactions[i])
		_, err := result.LastInsertId()
		if err != nil {
			log.WithError(err).WithField("migration", testTransactions[i]).Fatal("Can not execute migration")
		}
	}
}
