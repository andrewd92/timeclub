package order_service

type OrderService interface {
	Create(visits []int64) (interface{}, error)
}

type OrderServiceImpl struct {
}

func (o *OrderServiceImpl) Create(_ []int64) (interface{}, error) {
	return map[string]string{"response": "ok"}, nil
}
