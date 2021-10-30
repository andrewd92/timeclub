package order

import (
	"github.com/andrewd92/timeclub/visit_service/domain/visit"
	"time"
)

type Order struct {
	id            int64
	visit         *visit.Visit
	price         int64
	paymentMethod int64
	bonuses       int64
	opened        *time.Time
}
