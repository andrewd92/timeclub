package visit_service

import (
	"github.com/andrewd92/timeclub/visit_service/client/club_service"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	log "github.com/sirupsen/logrus"
	"time"
)

type VisitService interface {
	Create(clubId int64, cardId int64, dateTime time.Time) (interface{}, error)
	All(id int64, dateTime time.Time) ([]interface{}, error)
}

type visitServiceImpl struct {
	visitRepository   visit.Repository
	clubServiceClient club_service.ClubClient
	visitMarshaller   visit.Marshaller
	visitFactory      visit.Factory
}

func (s visitServiceImpl) All(clubId int64, dateTime time.Time) ([]interface{}, error) {
	visits, err := s.visitRepository.GetAll()
	if err != nil {
		return nil, err
	}

	club, clubServiceErr := s.clubServiceClient.GetById(clubId)
	if nil != clubServiceErr {
		return nil, clubServiceErr
	}

	response, responseErr := s.visitMarshaller.MarshalAll(visits, club, dateTime)

	if nil != responseErr {
		log.WithError(responseErr).Error("All visits response building error")
		return nil, responseErr
	}

	return response, nil
}

func (s visitServiceImpl) Create(clubId int64, cardId int64, dateTime time.Time) (interface{}, error) {
	newVisit := s.visitFactory.Create(clubId, cardId)

	visitModel, saveErr := s.visitRepository.Save(newVisit)
	if nil != saveErr {
		return nil, saveErr
	}

	club, clubServiceErr := s.clubServiceClient.GetById(clubId)
	if nil != clubServiceErr {
		return nil, clubServiceErr
	}

	marshal, marshalErr := s.visitMarshaller.Marshal(visitModel, dateTime, club)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return marshal, nil
}
