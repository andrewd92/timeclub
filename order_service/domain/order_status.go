package domain

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

type OrderStatus string

const (
	None   OrderStatus = "None"
	Open   OrderStatus = "OPEN"
	Cancel OrderStatus = "CANCEL"
	Paid   OrderStatus = "PAID"
)

func FromString(status string) (OrderStatus, error) {
	switch status {
	case string(Open):
		return Open, nil
	case string(Cancel):
		return Cancel, nil
	case string(Paid):
		return Paid, nil
	default:
		log.WithField("order_status", status).Error("can not parse order status")
		return None, errors.New("can not parse order status")
	}
}
