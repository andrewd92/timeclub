package event

type Repository interface {
	Save(event *Event) (*Event, error)
	GetAll() []*Event
}
