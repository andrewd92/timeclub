package card

type Repository interface {
	GetById(id int64) (*Card, error)
	Save(card *Card) (*Card, error)
}
