package visit

import (
	"github.com/andrewd92/timeclub/club_service/domain/client"
	"github.com/andrewd92/timeclub/club_service/domain/club"
	"github.com/andrewd92/timeclub/club_service/domain/order_details"
	"time"
)

func DefaultVisit() *Visit {
	now := time.Now()
	return DefaultVisitFrom(&now)
}

func DefaultVisitFrom(start *time.Time) *Visit {
	return NewVisit(
		1,
		start,
		client.DefaultClient(),
		club.DefaultClub(),
		order_details.DefaultOrderDetails(),
		"",
	)
}
