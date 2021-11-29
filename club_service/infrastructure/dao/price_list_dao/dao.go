package price_list_dao

import (
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	log "github.com/sirupsen/logrus"
)

const (
	insertList = `INSERT INTO price_list(name) values (:name);`
)

type PriceListDao struct {
	connection connection.Connection
}

func NewPriceListDao(connection connection.Connection) *PriceListDao {
	return &PriceListDao{connection: connection}
}

func (d PriceListDao) Insert(name string) (int64, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return 0, connectionErr
	}

	result, err := db.NamedExec(insertList, map[string]interface{}{
		"name": name,
	})

	if err != nil {
		log.WithError(err).Error("can not insert price list")
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.WithError(err).Error("can not insert price list")
		return 0, err
	}

	return insertId, nil
}
