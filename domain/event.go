package domain

import (
	"github.com/monicaribeiro/event-manager/dto"
	"github.com/monicaribeiro/event-manager/errs"
	"time"
)

type Event struct {
	Id        int64 `pg:"event_id,pk"`
	Name      string
	City      string
	State     string
	PhotoUrl  string `pg:"photo_url"`
	Datetime  time.Time
	CreatedOn time.Time `pg:"created_on"`
}

func (e Event) ToDto() dto.EventResponse {
	return dto.EventResponse{
		Id:        e.Id,
		Name:      e.Name,
		City:      e.City,
		State:     e.State,
		PhotoUrl:  e.PhotoUrl,
		Datetime:  e.Datetime,
		CreatedOn: e.CreatedOn,
	}
}

func NewEvent(name, city, state, photoUrl string, datetime time.Time) Event {
	return Event{
		Name:      name,
		City:      city,
		State:     state,
		PhotoUrl:  photoUrl,
		Datetime:  datetime,
		CreatedOn: time.Now(),
	}
}

type EventRepository interface {
	FindAll() ([]Event, *errs.AppError)
	Delete(int64) *errs.AppError
	ById(int64) (*Event, *errs.AppError)
	Create(*Event) *errs.AppError
}
