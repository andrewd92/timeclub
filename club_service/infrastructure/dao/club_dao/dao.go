package club_dao

import (
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	"github.com/andrewd92/timeclub/club_service/infrastructure/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	selectAll  = "SELECT * FROM `club`"
	selectById = "SELECT * FROM `club` WHERE id = ?"
	insert     = "INSERT INTO club(name, open_time, price_list_id, currency_id) VALUES (:name, :open_time, :price_list_id, :currency_id);"
	update     = "UPDATE club SET name = :name, open_time = :open_time, price_list_id = :price_list_id, currency_id = :currency_id WHERE id = :id;"
)

type ClubDao interface {
	GetAll() ([]*model.ClubModel, error)
	GetById(id int64) (*model.ClubModel, error)
	Insert(model *model.ClubModel) (int64, error)
	Update(dbModel *model.ClubModel) error
}

type ClubDaoImpl struct {
	connection connection.Connection
}

func NewClubDao(connection connection.Connection) ClubDao {
	return &ClubDaoImpl{connection: connection}
}

func (d ClubDaoImpl) GetAll() ([]*model.ClubModel, error) {
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

	var models []*model.ClubModel

	selectErr := db.Select(&models, selectAll)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select Clubs from db")
		return nil, selectErr
	}

	return models, nil
}

func (d ClubDaoImpl) GetById(id int64) (*model.ClubModel, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return nil, connectionErr
	}

	var dbModel = &model.ClubModel{}

	selectErr := db.Get(dbModel, selectById, id)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select Club entry from db")
		return nil, selectErr
	}

	return dbModel, nil
}

func (d ClubDaoImpl) Insert(dbModel *model.ClubModel) (int64, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return 0, connectionErr
	}

	result, err := db.NamedExec(insert, dbModel)

	if err != nil {
		log.WithError(err).WithField("model", dbModel).Error("can not insert Club")
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.WithError(err).WithField("model", dbModel).Error("can not insert Club")
		return 0, err
	}

	return insertId, nil
}

func (d ClubDaoImpl) Update(dbModel *model.ClubModel) error {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return connectionErr
	}

	_, err := db.NamedExec(update, dbModel)

	if err != nil {
		log.WithError(err).WithField("model", dbModel).Error("can not insert Club")
		return err
	}

	return nil
}
