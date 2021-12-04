package test

import (
	"github.com/andrewd92/timeclub/club_service/api"
	"github.com/andrewd92/timeclub/visit_service/client/club_service"
	"github.com/andrewd92/timeclub/visit_service/client/club_service/mocks"
	"github.com/andrewd92/timeclub/visit_service/infrastructure/connection"
	"github.com/andrewd92/timeclub/visit_service/server"
	"os"
	"testing"
	"time"
)

var clubClient *mocks.ClubClient
var club api.Club

func TestMain(m *testing.M) {
	connection.SetTestEnvironment()
	_ = os.Setenv("VIPER_CONFIG_NAME", "test")

	clubClient = new(mocks.ClubClient)
	initClub()
	clubClient.On("GetById", int64(1)).Return(&club, nil)

	club_service.SetMockInstance(clubClient)

	go server.StartApplication()
	time.Sleep(50 * time.Millisecond)
	os.Exit(m.Run())
}

func initClub() {
	club = api.Club{
		Id:       1,
		Name:     "Test",
		OpenTime: "12:00",
		Currency: &api.Currency{
			Name:      "USD",
			ShortName: "$",
		},
		Prices: []*api.Price{{
			PricePeriod: &api.PricePeriod{
				From: 0,
				To:   360,
			},
			ValuePerMinute: 10,
		}},
	}
}
