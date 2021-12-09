package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var testTransactions = []string{
	`CREATE TABLE IF NOT EXISTS orders (
    	id INTEGER NOT NULL PRIMARY KEY,
		status VARCHAR(255) NOT NULL DEFAULT 'OPEN',
		start TEXT NOT NULL,
		end TEXT
);`,
	`CREATE TABLE IF NOT EXISTS order_visits (
      id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
      order_id INT UNSIGNED NOT NULL,
      visit_id INT UNSIGNED NOT NULL,
      index (order_id)
);`,
	`INSERT INTO orders(id, status, visit_id, start, end)
VALUES (1, 'OPEN', datetime('now', '-1 Hour'), null);`,
	`INSERT INTO order_visits(id, order_id, visit_id) 
VALUES 
	(1, 1, 1),
	(2, 1, 4);`,
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
