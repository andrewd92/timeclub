package test

import (
	"github.com/andrewd92/timeclub/club_service/api"
	"github.com/andrewd92/timeclub/visit_service/client/club_service"
	"github.com/andrewd92/timeclub/visit_service/client/club_service/mocks"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/connection"
	"github.com/andrewd92/timeclub/visit_service/server"
	"github.com/andrewd92/timeclub/visit_service/utils"
	"os"
	"testing"
	"time"
)

var clubClient *mocks.ClubClient
var club *api.Club

func TestMain(m *testing.M) {
	connection.SetTestEnvironment()
	_ = os.Setenv("VIPER_CONFIG_NAME", "test")

	clubClient = new(mocks.ClubClient)
	initClub()
	clubClient.On("GetById", int64(1)).Return(club, nil)

	club_service.SetMockInstance(clubClient)

	go server.StartApplication()
	time.Sleep(50 * time.Millisecond)
	os.Exit(m.Run())
}

func initClub() {
	club = utils.DefaultClub()
}
