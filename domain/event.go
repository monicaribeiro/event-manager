package domain

import (
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

type EventRepository interface {
	FindAll() ([]Event, *errs.AppError)
	Delete(int64) *errs.AppError
	ById(int64) (*Event, *errs.AppError)
}
