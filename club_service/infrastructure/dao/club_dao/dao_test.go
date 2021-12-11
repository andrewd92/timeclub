package club_dao

import (
	"fmt"
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	model2 "github.com/andrewd92/timeclub/club_service/infrastructure/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var batch = []string{`
CREATE TABLE IF NOT EXISTS club (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    open_time VARCHAR(255) NOT NULL,
    price_list_id INTEGER,
    currency_id INTEGER NOT NULL
);`,
	`INSERT INTO club(name, open_time, price_list_id, currency_id) values ('t1', '12:00', 1, 1);`,
	`INSERT INTO club(name, open_time, price_list_id, currency_id) values ('t2', '12:00', 1, 1);`,
	`INSERT INTO club(name, open_time, price_list_id, currency_id) values ('t3', '12:00', 1, 1);`,
	`INSERT INTO club(name, open_time, price_list_id, currency_id) values ('t4', '12:00', 1, 1);`,
}

var dao ClubDao

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

	dao = NewClubDao(con)

	code := m.Run()

	os.Exit(code)
}

func TestClubDao_GetById(t *testing.T) {
	var clubId int64 = 2

	model, err := dao.GetById(clubId)

	expected := &model2.ClubModel{
		Id:          clubId,
		Name:        "t2",
		OpenTime:    "12:00",
		PriceListId: 1,
		CurrencyId:  1,
	}

	assert.Nil(t, err)
	assert.NotNil(t, model)
	assert.Equal(t, expected, model)
}

func TestClubDao_GetAll(t *testing.T) {
	models, err := dao.GetAll()

	assert.Nil(t, err)
	assert.NotNil(t, models)
	assert.Equal(t, 4, len(models))
	assert.Equal(t, "t3", models[2].Name)
}

func TestClubDaoImpl_Insert(t *testing.T) {
	dbModel := &model2.ClubModel{
		Name:        "t2",
		OpenTime:    "12:00",
		PriceListId: 1,
		CurrencyId:  1,
	}

	id, err := dao.Insert(dbModel)
	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.Greater(t, id, int64(0))
}

func TestClubDaoImpl_Update(t *testing.T) {
	expected := "new test name"

	dbModel := &model2.ClubModel{
		Name:        "t2",
		OpenTime:    "12:00",
		PriceListId: 1,
		CurrencyId:  1,
	}

	id, _ := dao.Insert(dbModel)
	dbModel.Id = id
	dbModel.Name = expected

	err := dao.Update(dbModel)
	assert.Nil(t, err)

	modelFromDb, err := dao.GetById(id)
	assert.Nil(t, err)
	assert.Equal(t, expected, modelFromDb.Name)
}
