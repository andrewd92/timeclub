package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var transactions = []string{
	`CREATE TABLE IF NOT EXISTS club (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_time VARCHAR(255) NOT NULL,
    price_list_id INT UNSIGNED,
    currency_id INT UNSIGNED NOT NULL
);`,
	`CREATE TABLE IF NOT EXISTS price (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    price_list_id INT UNSIGNED NOT NULL,
    period_from INT UNSIGNED NOT NULL,
    period_to INT UNSIGNED NOT NULL,
    value_per_minute DECIMAL(13,2),
    INDEX (price_list_id)
);`,
	`CREATE TABLE IF NOT EXISTS currency (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    short_name VARCHAR(255) NOT NULL
);`,
	`INSERT INTO currency(id, name, short_name) VALUES (1, 'US Dollar', 'USD')
ON DUPLICATE KEY UPDATE name = 'US Dollar', short_name = 'USD'`,
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
