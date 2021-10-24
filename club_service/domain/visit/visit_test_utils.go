package visit

import (
	"github.com/andrewd92/timeclub/club_service/domain/client"
	clubPkg "github.com/andrewd92/timeclub/club_service/domain/club"
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
		clubPkg.DefaultClub(),
		order_details.DefaultOrderDetails(),
		"",
	)
}

func VisitWithClub(club *clubPkg.Club) *Visit {
	now := time.Now()
	return NewVisit(
		1,
		&now,
		client.DefaultClient(),
		club,
		order_details.DefaultOrderDetails(),
		"",
	)
}
