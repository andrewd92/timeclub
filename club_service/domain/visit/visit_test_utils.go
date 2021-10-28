package visit

import (
	"github.com/andrewd92/timeclub/club_service/domain/card"
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
	return visit(start, clubPkg.DefaultClub())
}

func DefaultVisitWithClub(club *clubPkg.Club) *Visit {
	now := time.Now()
	return visit(&now, club)
}

func visit(start *time.Time, club *clubPkg.Club) *Visit {
	return NewVisit(
		1,
		start,
		club,
		order_details.DefaultOrderDetails(),
		"",
		card.DefaultCard(),
		client.DefaultClient().Name(),
	)
}
