package domain

type EventRepositoryStub struct {
	events []Event
}

func (s EventRepositoryStub) FindAll() ([]Event, error) {
	return s.events, nil
}

func NewEventRepositoryStub() EventRepositoryStub {
	events := []Event{
		{1001, "Meetup Devops", "Uberlândia", "MG", "https://raw.githubusercontent.com/ZupIT/charlescd/master/images/logo.png", "1621541106"},
		{1002, "Meetup QA", "São Paulo", "SP", "https://raw.githubusercontent.com/ZupIT/charlescd/master/images/logo.png", "1621541106"},
	}

	return EventRepositoryStub{events}
}
