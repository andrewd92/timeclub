package visit_service

import (
	"errors"
	clubClientMocks "github.com/andrewd92/timeclub/visit_service/client/club_service/mocks"
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	visitRepositoryMocks "github.com/andrewd92/timeclub/visit_service/domain/visit/mocks"
	"github.com/andrewd92/timeclub/visit_service/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestVisitServiceImpl_All(t *testing.T) {
	testingService, clubServiceClient, visitRepository := initService()
	mockDefaultClub(clubServiceClient)
	club := utils.DefaultClub()
	now := time.Now()

	visitEntity := visit.DefaultVisit()
	visits := []*visit.Visit{visitEntity}
	visitRepository.On("GetAll").Return(visits, nil)

	expected, _ := visit.MarshalAll(visits, club, now)
	actual, err := testingService.All(club.GetId(), now)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_AllReturnErrorWhenRepositoryError(t *testing.T) {
	testingService, clubServiceClient, visitRepository := initService()
	mockDefaultClub(clubServiceClient)
	club := utils.DefaultClub()
	now := time.Now()
	errMessage := "test err"

	visitRepository.On("GetAll").Return(nil, errors.New(errMessage))

	response, err := testingService.All(club.GetId(), now)

	assert.Nil(t, response)
	assert.Equal(t, errMessage, err.Error())
}

func Test_AllReturnErrorWhenClubClientError(t *testing.T) {
	testingService, clubServiceClient, visitRepository := initService()

	clubId := int64(3)
	now := time.Now()
	errMessage := "test err"

	visitEntity := visit.DefaultVisit()
	visits := []*visit.Visit{visitEntity}
	visitRepository.On("GetAll").Return(visits, nil)

	clubServiceClient.On("GetById", clubId).Return(nil, errors.New(errMessage))

	response, err := testingService.All(clubId, now)

	assert.Nil(t, response)
	assert.Equal(t, errMessage, err.Error())
}

func initService() (VisitService, *clubClientMocks.ClubClient, *visitRepositoryMocks.Repository) {
	visitRepository := new(visitRepositoryMocks.Repository)
	clubServiceClient := new(clubClientMocks.ClubClient)

	testingService := visitServiceImpl{
		visitRepository:   visitRepository,
		clubServiceClient: clubServiceClient,
	}

	return testingService, clubServiceClient, visitRepository
}

func mockDefaultClub(clubServiceClient *clubClientMocks.ClubClient) {
	club := utils.DefaultClub()
	clubServiceClient.On("GetById", club.GetId()).Return(club, nil)
}
