package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"strings"
)

var testTransactions = []string{
	`INSERT INTO orders(id, status, visit_id, start, end)
VALUES (1, 'OPEN', datetime('now', '-1 Hour'), null);`,
	`INSERT INTO order_visits(id, order_id, visit_id) 
VALUES 
	(1, 1, 1),
	(2, 1, 4);`,
}

func initTestMigration(db *sqlx.DB) {
	convertQueries()
	initMigration(db)

	for i := 0; i < len(testTransactions); i++ {
		log.WithField("migration", testTransactions[i]).Info("Run migration")
		result := db.MustExec(testTransactions[i])
		_, err := result.LastInsertId()
		if err != nil {
			log.WithError(err).WithField("migration", testTransactions[i]).Fatal("Can not execute migration")
		}
	}
}

func convertQueries() {
	for i, transaction := range transactions {

		if strings.Contains(transaction, "ADD INDEX") {
			transactions = append(transactions[:i], transactions[i+1:]...)
			continue
		}

		transaction = strings.Replace(transaction, " INT ", " INTEGER ", -1)
		transaction = strings.Replace(transaction, "AUTO_INCREMENT", "AUTOINCREMENT", -1)
		transaction = strings.Replace(transaction, "DATETIME", "TEXT", -1)
		transaction = strings.Replace(transaction, "UNSIGNED", "", -1)

		transactions[i] = transaction
	}
}
