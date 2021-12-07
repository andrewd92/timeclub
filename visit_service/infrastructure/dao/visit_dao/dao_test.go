package visit_dao

import (
	"fmt"
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/connection"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var batch = []string{`
CREATE TABLE IF NOT EXISTS visit (
    id INTEGER NOT NULL PRIMARY KEY,
    start TEXT NOT NULL,
    club_id INTEGER NOT NULL,
    order_details TEXT NOT NULL,
    comment VARCHAR(255) NOT NULL DEFAULT '',
    card_id INTEGER NOT NULL,
    client_name VARCHAR(255) NOT NULL
);`,
	`
INSERT INTO visit(id, start, club_id, order_details, comment, card_id, client_name)
VALUES
   (1, datetime('now', '-1 Hour'), 1, '[]', '', 1, 'Sasha'),
   (2, datetime('now', '-30 Minute'), 1, '[]', '', 2, 'Pasha'),
   (3, datetime('now', '-10 Minute'), 1, '[]', '', 3, 'Masha');`,
}

var dao VisitDao

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

	dao = NewVisitSqlDao(con)

	code := m.Run()

	os.Exit(code)
}

func Test_Get(t *testing.T) {
	visits, err := dao.GetAll()
	assert.Nil(t, err)
	assert.NotNil(t, visits)
	assert.True(t, len(visits) >= 3)
}

func Test_Insert(t *testing.T) {
	now := time.Now()
	id, err := dao.Insert(visit.NewVisit(
		&now,
		int64(1),
		order_details.DefaultOrderDetails(),
		"",
		int64(5),
		"Liza",
	))

	assert.Nil(t, err)
	assert.True(t, id > 3)
}
