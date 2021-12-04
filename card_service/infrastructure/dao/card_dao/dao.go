package card_dao

import (
	"database/sql"
	"github.com/andrewd92/timeclub/card_service/domain/card"
	"github.com/andrewd92/timeclub/card_service/infrastructure/connection"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	selectAll  = "SELECT * FROM card"
	selectById = "SELECT * FROM card WHERE id = ?"
	insertCard = "INSERT INTO card(discount, name, club_id) VALUES (:discount, :name, :clubId);"
)

type CardModel struct {
	Id       int64         `db:"id"`
	Discount float32       `db:"discount"`
	Name     string        `db:"name"`
	ClubId   sql.NullInt64 `db:"club_id"`
}

type CardDao interface {
	GetAll() ([]CardModel, error)
	GetById(id int64) (*CardModel, error)
	Insert(card *card.Card) (int64, error)
}

type CardSqlDao struct {
	connection connection.Connection
}

func NewCardSqlDao(connection connection.Connection) *CardSqlDao {
	return &CardSqlDao{connection: connection}
}

func (d CardSqlDao) GetById(id int64) (*CardModel, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return nil, connectionErr
	}
	defer closeDb(db)

	var model = &CardModel{}
	selectErr := db.Get(model, selectById, id)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectById).Error("Can not select card from db")
		return nil, selectErr
	}

	return model, nil
}

func (d CardSqlDao) GetAll() ([]CardModel, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return nil, connectionErr
	}
	defer closeDb(db)

	var models []CardModel

	selectErr := db.Select(&models, selectAll)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select cards from db")
		return nil, selectErr
	}

	return models, nil
}

func (d CardSqlDao) Insert(card *card.Card) (int64, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return 0, connectionErr
	}
	defer closeDb(db)

	result, err := db.NamedExec(insertCard, map[string]interface{}{
		"discount": card.Discount(),
		"name":     card.Name(),
		"clubId":   card.ClubId(),
	})

	if err != nil {
		log.WithError(err).WithField("card", card).Error("can not insert card")
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.WithError(err).WithField("card", card).Error("can not insert card")
		return 0, err
	}

	return insertId, nil
}

func closeDb(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		log.WithError(err).Error("Can not close sql connection")
	}
}
