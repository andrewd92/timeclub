package migration

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var testTransactions = []string{
	`CREATE TABLE IF NOT EXISTS club (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    open_time VARCHAR(255) NOT NULL,
    price_list_id INTEGER,
    currency_id INTEGER NOT NULL
);`,
	`CREATE TABLE IF NOT EXISTS price (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    price_list_id INTEGER NOT NULL,
    period_from INTEGER NOT NULL,
    period_to INTEGER NOT NULL,
    value_per_minute DOUBLE
);`,
	`CREATE TABLE IF NOT EXISTS currency (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    short_name VARCHAR(255) NOT NULL
);`,
	`CREATE TABLE IF NOT EXISTS price_list (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL
);`,
	`INSERT INTO currency(id, name, short_name) VALUES (1, 'US Dollar', 'USD')`,
	`INSERT INTO club(name, open_time, price_list_id, currency_id) values ('t1', '12:00', 1, 1);`,
	`INSERT INTO price_list(id, name) values (1, 'test price list');`,
	`INSERT INTO price(price_list_id, period_from, period_to, value_per_minute)
VALUES (1, 0, 360, 10);`,
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
