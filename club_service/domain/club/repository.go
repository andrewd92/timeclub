package club

type Repository interface {
	GetAll() ([]*Club, error)
	GetById(id int64) (*Club, error)

	Save(club *Club) (*Club, error)
}
