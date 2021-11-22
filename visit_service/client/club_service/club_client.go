package club_service

import (
	"github.com/andrewd92/timeclub/club_service/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

func GetById(id int64) (*api.Club, error) {
	url := viper.GetString("client.club.grpc.url")

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.WithError(err).Error("Fail to dial club server")
		return nil, err
	}
	defer conn.Close()

	client := api.NewClubServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	club, err := client.GetById(ctx, &api.Request{Id: id})

	if nil != err {
		log.WithError(err).WithField("id", id).Error("Fail to get club by id")
		return nil, err
	}

	return club, nil
}
