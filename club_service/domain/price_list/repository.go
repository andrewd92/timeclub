package price_list

type Repository interface {
	GetById(id int64) (*PriceList, error)
}
