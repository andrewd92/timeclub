package migration

import (
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS club (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_time VARCHAR(255) NOT NULL,
    price_list_id INT UNSIGNED,
    currency_id INT UNSIGNED NOT NULL
);
CREATE TABLE IF NOT EXISTS price (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    price_list_id INT UNSIGNED NOT NULL,
    period_from INT UNSIGNED NOT NULL,
    period_to INT UNSIGNED NOT NULL,
    value_per_minute DECIMAL(13,2)
);
`

func initMigration(db *sqlx.DB) {
	db.MustExec(schema)
}
