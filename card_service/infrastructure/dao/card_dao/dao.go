package card_dao

import (
	"database/sql"
	"github.com/andrewd92/timeclub/card_service/infrastructure/connection"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	selectAll = "SELECT * FROM card"
)

type CardModel struct {
	Id       int64         `db:"id"`
	Discount float32       `db:"discount"`
	Name     string        `db:"name"`
	ClubId   sql.NullInt64 `db:"club_id"`
}

type CardDao interface {
	GetAll() ([]CardModel, error)
}

type CardSqlDao struct {
	connection connection.Connection
}

func NewCardSqlDao(connection connection.Connection) *CardSqlDao {
	return &CardSqlDao{connection: connection}
}

func (d CardSqlDao) GetAll() ([]CardModel, error) {
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

	var models []CardModel

	selectErr := db.Select(&models, selectAll)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select cards from db")
		return nil, selectErr
	}

	return models, nil
}
