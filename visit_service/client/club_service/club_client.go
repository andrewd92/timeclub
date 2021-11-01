package club_service

import (
	"github.com/andrewd92/timeclub/club_service/api"
	configPkg "github.com/andrewd92/timeclub/visit_service/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GetById(id int64) (*api.Club, error) {
	config := configPkg.Instance()

	conn, err := grpc.Dial(config.Client.Club.Grpc.Url, grpc.WithInsecure())
	if err != nil {
		log.Printf("fail to dial: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := api.NewClubServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	club, err := client.GetById(ctx, &api.Request{Id: id})

	if nil != err {
		log.Printf("fail to get club by id #%d. Err: %v", id, err)
		return nil, err
	}

	return club, nil
}
