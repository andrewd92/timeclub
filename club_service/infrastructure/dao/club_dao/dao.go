package club_dao

import (
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	selectAll  = "SELECT * FROM club"
	selectById = "SELECT * FROM club WHERE id = ?"
)

type ClubModel struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	OpenTime    string `db:"open_time"`
	PriceListId int64  `db:"price_list_id"`
	CurrencyId  int64  `db:"currency_id"`
}

type ClubDao struct {
	connection connection.Connection
}

func NewClubDao(connection connection.Connection) *ClubDao {
	return &ClubDao{connection: connection}
}

func (d ClubDao) GetAll() ([]*ClubModel, error) {
	db, connectionErr := d.connection.Get()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.WithError(err).Error("Can not close sql connection")
		}
	}(db)
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
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return nil, connectionErr
	}

	var model = &ClubModel{}

	selectErr := db.Get(model, selectById, id)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select club entry from db")
		return nil, selectErr
	}

	return model, nil
}
