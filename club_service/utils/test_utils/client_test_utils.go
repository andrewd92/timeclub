package test_utils

import (
	"github.com/andrewd92/timeclub/club_service/domain"
	"time"
)

func DefaultClient() *domain.Client {

	birthday, _ := time.Parse("2006-01-02", "1990-01-01")

	return domain.NewClient(
		1,
		"Andy",
		"D",
		123321123,
		"andy@example.com",
		birthday,
		"abc123.jpg",
		0,
		DefaultClub(),
		"San Francisco",
		"",
		time.Now(),
		0,
	)
}
