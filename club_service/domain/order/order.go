package order

import (
	visit2 "github.com/andrewd92/timeclub/club_service/domain/visit"
	"time"
)

type Order struct {
	id            int64
	visit         *visit2.Visit
	price         int64
	paymentMethod int64
	bonuses       int64
	opened        *time.Time
}
