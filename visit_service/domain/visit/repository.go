package visit

type Repository interface {
	Save(visit *Visit) (*Visit, error)
	GetAll() []*Visit
}
