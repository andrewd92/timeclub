package event_service

import (
	"errors"
	"fmt"
	discountPkg "github.com/andrewd92/timeclub/club_service/domain/discount"
	eventPkg "github.com/andrewd92/timeclub/club_service/domain/event"
	"github.com/andrewd92/timeclub/club_service/domain/time_period"
	"github.com/andrewd92/timeclub/club_service/infrastructure/repository/event_repository"
	"time"
)

type eventServiceImpl struct {
	eventRepository eventPkg.Repository
}

type EventService interface {
	Create(name string, tag string, discount float32, start string, end string) (interface{}, error)
}

var service EventService

func Instance() EventService {
	if nil == service {
		service = &eventServiceImpl{
			eventRepository: event_repository.Instance(),
		}
	}

	return service
}

func (s eventServiceImpl) Create(name string, tag string, discount float32, start string, end string) (interface{}, error) {
	startTime, parseStartTimeErr := time.Parse("2006-01-02 15:04:05", start)
	if parseStartTimeErr != nil {
		msg := fmt.Sprintf("Invalid start time format: %s. Expected: 2006-01-02 15:04:05", start)
		return nil, errors.New(msg)
	}
	endTime, parseEndTimeErr := time.Parse("2006-01-02 15:04:05", end)
	if parseEndTimeErr != nil {
		msg := fmt.Sprintf("Invalid end time format: %s. Expected: 2006-01-02 15:04:05", end)
		return nil, errors.New(msg)
	}

	event := eventPkg.NewEvent(name, tag, *discountPkg.NewDiscount(discount), *time_period.NewTimePeriod(startTime, endTime))

	savedEvent, saveErr := s.eventRepository.Save(event)

	if saveErr != nil {
		return nil, saveErr
	}

	return savedEvent.Marshal(), nil
}
