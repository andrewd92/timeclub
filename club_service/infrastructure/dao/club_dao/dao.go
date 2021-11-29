package club_dao

//go:generate mockgen -destination=../../../utils/mocks/mock_club_dao.go -package=mocks github.com/andrewd92/timeclub/club_service/infrastructure/dao/club_dao ClubDao

import (
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	//selectAll  = "SELECT * FROM club"
	selectAll  = "SELECT club.id, club.name, club.open_time, club.price_list_id, club.currency_id FROM club"
	selectById = "SELECT * FROM club WHERE id = ?"
	insertClub = "INSERT INTO club(name, open_time, price_list_id, currency_id) VALUES (:name, :open_time, :price_list_id, :currency_id);"
)

type ClubModel struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	OpenTime    string `db:"open_time"`
	PriceListId int64  `db:"price_list_id"`
	CurrencyId  int64  `db:"currency_id"`
}

type ClubDao interface {
	GetAll() ([]*ClubModel, error)
	GetById(id int64) (*ClubModel, error)
	Insert(club *club.Club) (int64, error)
}

type ClubDaoImpl struct {
	connection connection.Connection
}

func NewClubDao(connection connection.Connection) ClubDao {
	return &ClubDaoImpl{connection: connection}
}

func (d ClubDaoImpl) GetAll() ([]*ClubModel, error) {
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

func (d ClubDaoImpl) GetById(id int64) (*ClubModel, error) {
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

func (d ClubDaoImpl) Insert(club *club.Club) (int64, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return 0, connectionErr
	}

	result, err := db.NamedExec(insertClub, map[string]interface{}{
		"name":          club.Name(),
		"open_time":     club.OpenTime(),
		"price_list_id": 1,
		"currency_id":   club.Currency().Id(),
	})

	if err != nil {
		log.WithError(err).WithField("club", club).Error("can not insert club")
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.WithError(err).WithField("club", club).Error("can not insert club")
		return 0, err
	}

	return insertId, nil
}
