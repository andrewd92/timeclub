package migration

import (
	"github.com/jmoiron/sqlx"
)

const club = `CREATE TABLE IF NOT EXISTS club (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_time VARCHAR(255) NOT NULL,
    price_list_id INT UNSIGNED,
    currency_id INT UNSIGNED NOT NULL
);
`
const price = `CREATE TABLE IF NOT EXISTS price (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    price_list_id INT UNSIGNED NOT NULL,
    period_from INT UNSIGNED NOT NULL,
    period_to INT UNSIGNED NOT NULL,
    value_per_minute DECIMAL(13,2),
    INDEX (price_list_id)
);
`
const currency = `CREATE TABLE IF NOT EXISTS currency (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    short_name VARCHAR(255) NOT NULL
);
`
const USD = `INSERT INTO currency(id, name, short_name) VALUES (1, 'US Dollar', 'USD')
ON DUPLICATE KEY UPDATE name = 'US Dollar', short_name = 'USD'`

func initMigration(db *sqlx.DB) {
	db.MustExec(club)
	db.MustExec(price)
	db.MustExec(currency)
	db.MustExec(USD)
}
