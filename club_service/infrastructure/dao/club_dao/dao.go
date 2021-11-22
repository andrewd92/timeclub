package club_dao

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	selectAll  = "SELECT * FROM club"
	selectById = "SELECT * FROM club WHERE id = $1"
)

type ClubModel struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	OpenTime    string `db:"open_time"`
	PriceListId int64  `db:"price_list_id"`
	CurrencyId  int64  `db:"currency_id"`
}

type ClubDao struct {
}

func (d ClubDao) GetAll() ([]*ClubModel, error) {
	db, connectionErr := getConnection()
	if connectionErr != nil {
		return nil, connectionErr
	}

	var models []*ClubModel

	selectErr := db.Select(&models, selectAll)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select clubs from db")
		return nil, selectErr
	}

	return models, nil
}

func (d ClubDao) GetById(id int64) (*ClubModel, error) {
	db, connectionErr := getConnection()
	if connectionErr != nil {
		return nil, connectionErr
	}

	var model *ClubModel

	selectErr := db.Get(model, selectById, id)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select club entry from db")
		return nil, selectErr
	}

	return model, nil
}

func getConnection() (*sqlx.DB, error) {
	dbUrl := viper.GetString("db.url")

	db, err := sqlx.Connect("mysql", dbUrl)

	if err != nil {
		log.WithError(err).Error("Can not connect to db.")
		return nil, err
	}

	return db, nil
}
