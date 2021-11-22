package club

type Repository interface {
	GetAll() ([]*Club, error)
	GetById(id int64) (*Club, error)
}
