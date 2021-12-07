package visit_service

import (
	"errors"
	"github.com/andrewd92/timeclub/visit_service/client/club_service"
	clubClientMocks "github.com/andrewd92/timeclub/visit_service/client/club_service/mocks"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	visitDomainMocks "github.com/andrewd92/timeclub/visit_service/domain/visit/mocks"
	"github.com/andrewd92/timeclub/visit_service/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const cardId = int64(3)

var defaultClub = utils.DefaultClub()
var now = time.Now()

func TestVisitServiceImpl_All(t *testing.T) {
	clubServiceClient, repository, _, factory := mocks()
	testingService := initServiceWithDefaultMarshaller(clubServiceClient, repository, factory)
	mockDefaultClub(clubServiceClient)

	visitEntity := visit.DefaultVisit()
	visits := []*visit.Visit{visitEntity}
	repository.On("GetAll").Return(visits, nil)

	expected, _ := (visit.MarshallerImpl{}).MarshalAll(visits, defaultClub, now)
	actual, err := testingService.All(defaultClub.GetId(), now)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_AllReturnErrorWhenRepositoryError(t *testing.T) {
	clubServiceClient, repository, marshaller, factory := mocks()
	testingService := initTestingService(clubServiceClient, repository, marshaller, factory)
	mockDefaultClub(clubServiceClient)
	errMessage := "test err"

	repository.On("GetAll").Return(nil, errors.New(errMessage))

	response, err := testingService.All(defaultClub.GetId(), now)

	assert.Nil(t, response)
	assert.Equal(t, errMessage, err.Error())
}

func Test_AllReturnErrorWhenClubClientError(t *testing.T) {
	clubServiceClient, repository, marshaller, factory := mocks()
	testingService := initTestingService(clubServiceClient, repository, marshaller, factory)

	clubId := int64(3)
	errMessage := "test err"

	visitEntity := visit.DefaultVisit()
	visits := []*visit.Visit{visitEntity}
	repository.On("GetAll").Return(visits, nil)

	clubServiceClient.On("GetById", clubId).Return(nil, errors.New(errMessage))

	response, err := testingService.All(clubId, now)

	assert.Nil(t, response)
	assert.Equal(t, errMessage, err.Error())
}

func Test_AllReturnErrorWhenMarshallingError(t *testing.T) {
	clubServiceClient, repository, marshaller, factory := mocks()
	testingService := initTestingService(clubServiceClient, repository, marshaller, factory)
	mockDefaultClub(clubServiceClient)
	errMessage := "test err"

	visitEntity := visit.DefaultVisit()
	visits := []*visit.Visit{visitEntity}
	repository.On("GetAll").Return(visits, nil)

	marshaller.On("MarshalAll", visits, defaultClub, now).Return(nil, errors.New(errMessage))

	response, err := testingService.All(defaultClub.Id, now)

	assert.Nil(t, response)
	assert.Equal(t, errMessage, err.Error())
}

func Test_Create(t *testing.T) {
	clubServiceClient, repository, _, factory := mocks()
	testingService := initServiceWithDefaultMarshaller(clubServiceClient, repository, factory)
	visitEntity := visit.DefaultVisit()

	factory.On("Create", defaultClub.Id, cardId).Return(visitEntity)
	repository.On("Save", visitEntity).Return(visitEntity, nil)
	mockDefaultClub(clubServiceClient)

	expected, _ := (visit.MarshallerImpl{}).Marshal(visitEntity, now, defaultClub)
	actual, err := testingService.Create(defaultClub.GetId(), cardId, now)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_CreateReturnErrorWhenRepositoryError(t *testing.T) {
	clubServiceClient, repository, marshaller, factory := mocks()
	testingService := initTestingService(clubServiceClient, repository, marshaller, factory)
	errMessage := "test err"
	visitEntity := visit.DefaultVisit()

	factory.On("Create", defaultClub.Id, cardId).Return(visitEntity)
	repository.On("Save", visitEntity).Return(nil, errors.New(errMessage))

	response, err := testingService.Create(defaultClub.GetId(), cardId, now)

	assert.Nil(t, response)
	assert.Equal(t, errMessage, err.Error())
}

func Test_CreateReturnErrorWhenClubServiceError(t *testing.T) {
	clubServiceClient, repository, marshaller, factory := mocks()
	testingService := initTestingService(clubServiceClient, repository, marshaller, factory)
	errMessage := "test err"
	visitEntity := visit.DefaultVisit()

	factory.On("Create", defaultClub.Id, cardId).Return(visitEntity)
	repository.On("Save", visitEntity).Return(visitEntity, nil)
	clubServiceClient.On("GetById", defaultClub.Id).Return(nil, errors.New(errMessage))

	response, err := testingService.Create(defaultClub.GetId(), cardId, now)

	assert.Nil(t, response)
	assert.Equal(t, errMessage, err.Error())
}

func Test_CreateReturnErrorWhenMarshallingError(t *testing.T) {
	clubServiceClient, repository, marshaller, factory := mocks()
	testingService := initTestingService(clubServiceClient, repository, marshaller, factory)
	errMessage := "test err"
	visitEntity := visit.DefaultVisit()

	factory.On("Create", defaultClub.Id, cardId).Return(visitEntity)
	repository.On("Save", visitEntity).Return(visitEntity, nil)
	mockDefaultClub(clubServiceClient)
	marshaller.On("Marshal", visitEntity, now, defaultClub).Return(nil, errors.New(errMessage))

	response, err := testingService.Create(defaultClub.GetId(), cardId, now)

	assert.Nil(t, response)
	assert.Equal(t, errMessage, err.Error())
}

func initServiceWithDefaultMarshaller(
	clubServiceClient club_service.ClubClient,
	visitRepository visit.Repository,
	visitFactory visit.Factory,
) VisitService {
	visitMarshaller := visit.MarshallerImpl{}

	return initTestingService(clubServiceClient, visitRepository, visitMarshaller, visitFactory)
}

func initTestingService(
	clubServiceClient club_service.ClubClient,
	visitRepository visit.Repository,
	visitMarshaller visit.Marshaller,
	visitFactory visit.Factory,
) VisitService {
	testingService := visitServiceImpl{
		visitRepository:   visitRepository,
		clubServiceClient: clubServiceClient,
		visitMarshaller:   visitMarshaller,
		visitFactory:      visitFactory,
	}

	return testingService
}

func mocks() (*clubClientMocks.ClubClient, *visitDomainMocks.Repository, *visitDomainMocks.Marshaller, *visitDomainMocks.Factory) {
	visitRepository := new(visitDomainMocks.Repository)
	clubServiceClient := new(clubClientMocks.ClubClient)
	visitMarshaller := new(visitDomainMocks.Marshaller)
	visitFactory := new(visitDomainMocks.Factory)

	return clubServiceClient, visitRepository, visitMarshaller, visitFactory
}

func mockDefaultClub(clubServiceClient *clubClientMocks.ClubClient) {
	club := utils.DefaultClub()
	clubServiceClient.On("GetById", club.GetId()).Return(club, nil)
}
