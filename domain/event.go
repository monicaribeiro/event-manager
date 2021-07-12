package domain

type Event struct {
	Id            int
	Name          string
	City          string
	State         string
	photoUrl      string
	eventDateTime string
}

type EventRepository interface {
	FindAll() ([]Event, error)
}
