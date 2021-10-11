package domain

import "time"

type Order struct {
	id            int64
	visit         *Visit
	price         int64
	paymentMethod int64
	bonuses       int64
	opened        *time.Time
}
