package card_dao

import (
	"database/sql"
	"fmt"
	"github.com/andrewd92/timeclub/card_service/infrastructure/connection"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var batch = []string{`
	CREATE TABLE IF NOT EXISTS card (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		discount DECIMAL(13,2) NOT NULL,
		name VARCHAR(255) NOT NULL,
		club_id INTEGER
	)`,
	`INSERT INTO card values (1, 0.0, 'Guest', null);`,
	`INSERT INTO card values (2, 5.0, 'Personal', 1);`,
	`INSERT INTO card values (3, 20.0, 'VIP', 1);`,
}

var dao CardDao

func TestMain(m *testing.M) {
	connection.SetTestEnvironment()
	con := connection.Instance()
	db, _ := con.Get()

	for _, query := range batch {
		db.MustExec(query)
	}

	func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("DB close Error: " + err.Error())
		}
	}(db)

	dao = NewCardSqlDao(con)

	code := m.Run()

	os.Exit(code)
}

func TestCardDao_GetById(t *testing.T) {
	var cardId int64 = 2

	model, err := dao.GetById(cardId)

	var expected = &CardModel{
		Id:       cardId,
		Discount: 5.0,
		Name:     "Personal",
		ClubId:   sql.NullInt64{Int64: 1, Valid: true},
	}
	assert.Nil(t, err)
	assert.NotNil(t, model)
	assert.Equal(t, expected, model)
}
