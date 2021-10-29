package card_template

type Repository interface {
	GetAll() []*CardTemplate
	GetById(id int64) (*CardTemplate, error)
	Save(cardTemplate *CardTemplate) (*CardTemplate, error)
}
