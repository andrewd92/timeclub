package currency_dao

import (
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	log "github.com/sirupsen/logrus"
)

const (
	selectById = "SELECT * FROM `currency` WHERE `id` = ?"
	selectAll  = "SELECT * FROM `currency`"
)

type CurrencyModel struct {
	Id        int64  `db:"id"`
	Name      string `db:"name"`
	ShortName string `db:"short_name"`
}

type CurrencyDao struct {
	connection connection.Connection
}

func NewCurrencyDao(connection connection.Connection) *CurrencyDao {
	return &CurrencyDao{connection: connection}
}

func (d CurrencyDao) GetAll() ([]*CurrencyModel, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return nil, connectionErr
	}

	var models []*CurrencyModel

	selectErr := db.Select(&models, selectAll)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select currencies from db")
		return nil, selectErr
	}

	return models, nil
}

func (d CurrencyDao) GetById(id int64) (*CurrencyModel, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return nil, connectionErr
	}

	var model *CurrencyModel

	selectErr := db.Get(&model, selectById, id)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectById).Error("Can not select currency entry from db")
		return nil, selectErr
	}

	return model, nil
}
