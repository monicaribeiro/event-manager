package domain

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

type EventRepositoryDb struct {
	db *pg.DB
}

func (d EventRepositoryDb) FindAll() ([]Event, error) {
	var events []Event
	err := d.db.Model(&events).Select()

	if err != nil {
		log.Println("Error while querying event table." + err.Error())
		return nil, err
	}
	return events, nil
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
