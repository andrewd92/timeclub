package visit_dao

import (
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/connection"
	"github.com/andrewd92/timeclub/visit_service/utils"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	selectAll   = "SELECT * FROM visit"
	insertVisit = `INSERT INTO visit(start, club_id, order_details, comment, card_id, client_name)
VALUES (:start, :club_id, :order_details, :comment, :card_id, :client_name)`
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
	Insert(visit *visit.Visit) (int64, error)
}

func Instance() VisitDao {
	return &VisitSqlDao{connection: connection.Instance()}
}

type VisitSqlDao struct {
	connection connection.Connection
}

func NewVisitSqlDao(connection connection.Connection) *VisitSqlDao {
	return &VisitSqlDao{connection: connection}
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

func (d VisitSqlDao) Insert(visit *visit.Visit) (int64, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return 0, connectionErr
	}

	startTimeInUtc := visit.Start().UTC()

	log.WithField("start_time", visit.Start().Format(utils.TimeFormat)).Info("Start time")
	log.WithField("start_time_utc", startTimeInUtc.Format(utils.TimeFormat)).Info("Start time UTC")

	params := map[string]interface{}{
		"start":         startTimeInUtc.Format(utils.TimeFormat),
		"club_id":       visit.ClubId(),
		"order_details": "[]",
		"comment":       visit.Comment(),
		"card_id":       visit.CardId(),
		"client_name":   visit.ClientName(),
	}
	result, err := db.NamedExec(insertVisit, params)

	if err != nil {
		log.WithError(err).WithField("visit", visit).WithField("params", params).
			Error("can not insert visit")
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.WithError(err).WithField("visit", visit).WithField("params", params).
			Error("can not get insert visit result")
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
