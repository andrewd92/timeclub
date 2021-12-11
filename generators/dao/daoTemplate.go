package dao

var templateStr = `
package {{ .Package }}

import (
	"github.com/andrewd92/timeclub/{{ .Service }}/infrastructure/connection"
	"github.com/andrewd92/timeclub/{{ .Service }}/infrastructure/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	selectAll  = "SELECT * FROM {{ .Table }}"
	selectById = "SELECT * FROM {{ .Table }} WHERE id = ?"
	insert = "INSERT INTO club({{ .Fields }}) VALUES ({{ .Placeholders }});"
	update = "UPDATE club SET {{ .UpdateFields }} WHERE id = :id;"
)

type {{ .Name }}Dao interface {
	GetAll() ([]*{{ .Model }}, error)
	GetById(id int64) (*{{ .Model }}, error)
	Insert(model *{{ .Model }}) (int64, error)
	Update(dbModel *{{ .Model }}) error
}

type {{ .Name }}DaoImpl struct {
	connection connection.Connection
}

func New{{ .Name }}Dao(connection connection.Connection) {{ .Name }}Dao {
	return &{{ .Name }}DaoImpl{connection: connection}
}

func (d {{ .Name }}DaoImpl) GetAll() ([]*{{ .Model }}, error) {
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

	var models []*{{ .Model }}

	selectErr := db.Select(&models, selectAll)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select {{ .Name }}s from db")
		return nil, selectErr
	}

	return models, nil
}

func (d {{ .Name }}DaoImpl) GetById(id int64) (*{{ .Model }}, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return nil, connectionErr
	}

	var dbModel = &{{ .Model }}{}

	selectErr := db.Get(dbModel, selectById, id)
	if selectErr != nil {
		log.WithError(selectErr).WithField("query", selectAll).Error("Can not select {{ .Name }} entry from db")
		return nil, selectErr
	}

	return dbModel, nil
}

func (d {{ .Name }}DaoImpl) Insert(dbModel *{{ .Model }}) (int64, error) {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return 0, connectionErr
	}

	result, err := db.NamedExec(insert, dbModel)

	if err != nil {
		log.WithError(err).WithField("model", dbModel).Error("can not insert {{ .Name }}")
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.WithError(err).WithField("model", dbModel).Error("can not insert {{ .Name }}")
		return 0, err
	}

	return insertId, nil
}


func (d {{ .Name }}DaoImpl) Update(dbModel *{{ .Model }}) error {
	db, connectionErr := d.connection.Get()
	if connectionErr != nil {
		return connectionErr
	}

	_, err := db.NamedExec(update, dbModel)

	if err != nil {
		log.WithError(err).WithField("model", dbModel).Error("can not insert {{ .Name }}")
		return err
	}

	return nil
}
`
