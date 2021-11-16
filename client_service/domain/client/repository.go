package client

type Repository interface {
	GetById(id int64) (*Client, error)
}
