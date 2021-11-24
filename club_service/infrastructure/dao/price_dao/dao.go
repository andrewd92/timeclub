package price_dao

import (
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	log "github.com/sirupsen/logrus"
)

const (
	selectByListId = "SELECT * FROM `price` WHERE `price_list_id` = ?"
)

type PriceModel struct {
	Id             int64   `db:"id"`
	PriceListId    int64   `db:"price_list_id"`
	PeriodFrom     int     `db:"period_from"`
	PeriodTo       int     `db:"period_to"`
	ValuePerMinute float32 `db:"value_per_minute"`
}

type PriceDao struct {
	connection connection.Connection
}

func (d PriceDao) GetByPriceListId(priceListId int64) ([]*PriceModel, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return nil, connectionErr
	}

	var models []*PriceModel

	selectErr := db.Select(&models, selectByListId, priceListId)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectByListId).Error("Can not select prices from db")
		return nil, selectErr
	}

	return models, nil
}
