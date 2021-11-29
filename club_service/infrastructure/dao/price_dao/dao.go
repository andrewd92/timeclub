package price_dao

import (
	"github.com/andrewd92/timeclub/club_service/domain/price_list/price"
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	log "github.com/sirupsen/logrus"
)

const (
	selectByListId = "SELECT * FROM price WHERE `price_list_id` = ?"
	deleteByListId = "DELETE FROM `price` WHERE `price_list_id` = :id"
	insertPrice    = `INSERT INTO price(price_list_id, period_from, period_to, value_per_minute) 
VALUES (:price_list_id, :period_from, :period_to, :value_per_minute);`
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

func NewPriceDao(connection connection.Connection) *PriceDao {
	return &PriceDao{connection: connection}
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

func (d PriceDao) Insert(price *price.Price, priceListId int64) (int64, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return 0, connectionErr
	}

	result, err := db.NamedExec(insertPrice, map[string]interface{}{
		"price_list_id":    priceListId,
		"period_from":      price.PricePeriod().From(),
		"period_to":        price.PricePeriod().To(),
		"value_per_minute": price.ValuePerMinute(),
	})

	if err != nil {
		log.WithError(err).WithField("price", price).Error("can not insert price")
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.WithError(err).WithField("price", price).Error("can not insert price")
		return 0, err
	}

	return insertId, nil
}

func (d PriceDao) DeletePriceList(priceListId int64) error {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return connectionErr
	}

	_, err := db.NamedExec(deleteByListId, map[string]interface{}{
		"id": priceListId,
	})

	if err != nil {
		log.WithError(err).WithField("price_list_id", priceListId).Error("can not delete prices")
		return err
	}

	return nil
}
