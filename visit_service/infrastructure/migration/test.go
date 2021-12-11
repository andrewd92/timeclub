package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"strings"
)

var testTransactions = []string{
	`INSERT INTO visit(id, start, club_id, order_details, comment, card_id, client_name)
VALUES
   (1, datetime('now', '-1 Hour'), 1, '[]', '', 1, 'Sasha');`,
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
