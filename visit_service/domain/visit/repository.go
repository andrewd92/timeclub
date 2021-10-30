package visit

type Repository interface {
	CreateVisit(clubId int64, cardId int64) (*Visit, error)
	GetAll() []*Visit
}
