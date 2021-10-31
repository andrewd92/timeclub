package visit

import (
	"github.com/andrewd92/timeclub/visit_service/domain/order_details"
	"time"
)

func DefaultVisit() *Visit {
	now := time.Now()
	return DefaultVisitFrom(&now)
}

func DefaultVisitFrom(start *time.Time) *Visit {
	return visit(start)
}

func visit(start *time.Time) *Visit {
	return NewVisit(
		start,
		1,
		order_details.DefaultOrderDetails(),
		"",
		1,
		"Andy",
	).WithId(1)
}
