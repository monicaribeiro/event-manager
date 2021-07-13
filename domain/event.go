package domain

import "time"

type Event struct {
	Id        int64 `pg:"event_id"`
	Name      string
	City      string
	State     string
	PhotoUrl  string `pg:"photo_url"`
	Datetime  time.Time
	CreatedOn time.Time `pg:"created_on"`
}

type EventRepository interface {
	FindAll() ([]Event, error)
}
