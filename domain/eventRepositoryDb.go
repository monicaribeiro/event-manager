package domain

import (
	"context"
	"github.com/monicaribeiro/event-manager/errs"
	"log"

	"github.com/go-pg/pg/v10"
)

type EventRepositoryDb struct {
	db *pg.DB
}

func (e EventRepositoryDb) FindAll() ([]Event, *errs.AppError) {
	var events []Event
	err := e.db.Model(&events).Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, errs.NewNotFoundError("Events not found")
		} else {
			log.Println("Error while querying event table." + err.Error())
			return nil, errs.NewUnexpectedErrorError("Unexpected database error")
		}
	}
	return events, nil
}

func (e EventRepositoryDb) Delete(id int64) *errs.AppError {
	event := &Event{Id: id}

	_, err := e.db.Model(event).WherePK().ForceDelete()

	if err != nil {
		log.Println("Error while deleting event." + err.Error())
		return errs.NewUnexpectedErrorError("Unexpected database error")
	}

	return nil
}

func (e EventRepositoryDb) ById(id int64) (*Event, *errs.AppError) {
	event := &Event{Id: id}

	err := e.db.Model(event).WherePK().Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, errs.NewNotFoundError("Event not found")
		} else {
			log.Println("Error while selecting event." + err.Error())
			log.Printf("Id: %d\n", id)
			return nil, errs.NewUnexpectedErrorError("Unexpected database error")
		}
	}
	return event, nil
}

func NewEventRepositoryDb() EventRepositoryDb {
	opt, err := pg.ParseURL("postgres://admin:admin@localhost:5432/event-manager?sslmode=disable")

	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	ctx := context.Background()
	_, err = db.ExecContext(ctx, "SELECT 1")

	if err != nil {
		panic(err)
	}

	return EventRepositoryDb{db}
}
