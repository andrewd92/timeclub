package price_list

//go:generate mockgen -destination=../../utils/mocks/mock_price_list_repository.go -package=mocks -mock_names=Repository=MockPriceListRepository . Repository

type Repository interface {
	GetById(id int64) (*PriceList, error)
	Save(list *PriceList) (*PriceList, error)
}
