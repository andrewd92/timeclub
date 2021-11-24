package currency

type Repository interface {
	GetAll() ([]*Currency, error)
	GetById(id int64) (*Currency, error)
}
