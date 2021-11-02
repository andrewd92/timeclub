package currency

type Repository interface {
	GetAll() []*Currency
	GetById(id int64) (*Currency, error)
}
