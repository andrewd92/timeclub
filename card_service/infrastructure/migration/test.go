package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var testTransactions = []string{
	`CREATE TABLE IF NOT EXISTS card (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    discount DECIMAL(13,2) NOT NULL,
    name VARCHAR(255) NOT NULL,
    club_id INTEGER
)`,
}

func initTestMigration(db *sqlx.DB) {
	log.Info("Run test migrations...")

	firstTransaction := 0
	for i := firstTransaction; i < len(testTransactions); i++ {
		log.WithField("migration", testTransactions[i]).Info("Run migration")
		result := db.MustExec(testTransactions[i])
		_, err := result.LastInsertId()
		if err != nil {
			log.WithError(err).WithField("migration", testTransactions[i]).Fatal("Can not execute migration")
		}
	}

	createTestGuestCards(db)
}

func createTestGuestCards(db *sqlx.DB) {
	row := db.QueryRow(selectGuestCardsCount)
	var cardsCount int
	err := row.Scan(&cardsCount)

	if err != nil {
		log.WithError(err).WithField("migration", selectGuestCardsCount).Fatal("Can not execute migration query")
	}

	if cardsCount != 0 {
		return
	}

	log.Info("Creating guest cards")
	for i := 0; i < 3; i++ {
		result := db.MustExec(insertGuestCard)
		_, err := result.LastInsertId()
		if err != nil {
			log.WithError(err).WithField("query", insertGuestCard).Error("Can create guest cards")
		}
	}
}
