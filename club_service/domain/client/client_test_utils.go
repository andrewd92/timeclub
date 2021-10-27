package client

import (
	"github.com/andrewd92/timeclub/club_service/domain/card"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"time"
)

func DefaultClient() *Client {

	birthday, _ := time.Parse("2006-01-02", "1990-01-01")

	return NewClient(
		1,
		"Andy",
		"D",
		123321123,
		"andy@example.com",
		birthday,
		"abc123.jpg",
		0,
		club.DefaultClub(),
		"San Francisco",
		"",
		time.Now(),
		0,
		card.DefaultCard(),
	)
}
