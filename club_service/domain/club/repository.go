package club

type Repository interface {
	GetAll() []*Club
	GetById(id int64) (*Club, error)
}
