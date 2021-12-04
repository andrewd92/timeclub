package visit_dao

import (
	"github.com/andrewd92/timeclub/visit_service/infrastructure/connection"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	selectAll = "SELECT * FROM visit"
)

type VisitModel struct {
	Id           int64  `db:"id"`
	Start        string `db:"start"`
	ClubId       int64  `db:"club_id"`
	OrderDetails string `db:"order_details"`
	Comment      string `db:"comment"`
	CardId       int64  `db:"card_id"`
	ClientName   string `db:"client_name"`
}

type VisitDao interface {
	GetAll() ([]VisitModel, error)
}

func Instance() VisitDao {
	return &VisitSqlDao{connection: connection.Instance()}
}

type VisitSqlDao struct {
	connection connection.Connection
}

func (d VisitSqlDao) GetAll() ([]VisitModel, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return nil, connectionErr
	}
	defer closeDb(db)

	var models []VisitModel

	selectErr := db.Select(&models, selectAll)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select cards from db")
		return nil, selectErr
	}

	return models, nil
}

func closeDb(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		log.WithError(err).Error("Can not close sql connection")
	}
}
