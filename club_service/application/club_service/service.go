package club_service

import (
	"errors"
	"github.com/andrewd92/timeclub/club_service/api/http/create"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/currency"
	"github.com/andrewd92/timeclub/club_service/domain/price_list"
	log "github.com/sirupsen/logrus"
)

type ClubService interface {
	Create(request create.Request) (interface{}, error)
}

type clubServiceImpl struct {
	repository club.Repository
}

func (s clubServiceImpl) Create(request create.Request) (interface{}, error) {
	clubEntity := club.NewClub(request.Name, request.OpenTime, price_list.Empty(), currency.USD())

	newClub, err := s.repository.Save(clubEntity)
	if err != nil {
		log.WithError(err).WithField("clubEntity", clubEntity).Error("can not store clubEntity")
		return nil, errors.New("db error")
	}

	return newClub.Marshal(), nil
}
