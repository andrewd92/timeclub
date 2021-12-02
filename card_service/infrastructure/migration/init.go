package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const selectGuestCardsCount = "SELECT COUNT(*) as card_count FROM card WHERE id <= 100;"
const insertGuestCard = "INSERT INTO card(discount, name, club_id) VALUES (0.0, 'Guest Card', null);"

var transactions = []string{
	`CREATE TABLE IF NOT EXISTS card (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    discount DECIMAL(13,2) NOT NULL,
    name VARCHAR(255) NOT NULL,
    club_id INT UNSIGNED,
    index idx_card_club(club_id)
)`,
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

	createGuestCards(db)
}

func createGuestCards(db *sqlx.DB) {
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
	for i := 0; i < 100; i++ {
		result := db.MustExec(insertGuestCard)
		_, err := result.LastInsertId()
		if err != nil {
			log.WithError(err).WithField("query", insertGuestCard).Error("Can create guest cards")
		}
	}
}
