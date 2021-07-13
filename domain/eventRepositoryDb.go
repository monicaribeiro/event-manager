package domain

import (
	"context"
	"fmt"
	"github.com/monicaribeiro/event-manager/errs"
	"github.com/monicaribeiro/event-manager/logger"
	"os"

	"github.com/go-pg/pg/v10"
)

type EventRepositoryDb struct {
	db *pg.DB
}

func (e EventRepositoryDb) FindAll(state string) ([]Event, *errs.AppError) {
	var events []Event
	var err error

	if state == "" {
		err = e.db.Model(&events).Select()
	} else {
		err = e.db.Model(&events). Where("state = ?", state).Select()
	}

	if err != nil {
		if err == pg.ErrNoRows {
			logger.Error("No event found.")
			return nil, errs.NewNotFoundError("Events not found")
		} else {
			logger.Error("Error while querying event table." + err.Error())
			return nil, errs.NewUnexpectedErrorError("Unexpected database error")
		}
	}
	return events, nil
}

func (e EventRepositoryDb) Delete(id int64) *errs.AppError {
	event := &Event{Id: id}

	_, err := e.db.Model(event).WherePK().ForceDelete()

	if err != nil {
		logger.Error("Error while deleting event." + err.Error())
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
			logger.Error("Error while selecting event." + err.Error())
			return nil, errs.NewUnexpectedErrorError("Unexpected database error")
		}
	}
	return event, nil
}

func (e EventRepositoryDb) Create(event *Event) *errs.AppError {

	_, err := e.db.Model(event).Insert()

	if err != nil {
		logger.Error("Error while creating event." + err.Error())
		return errs.NewUnexpectedErrorError("Unexpected database error")
	}

	return nil
}

func NewEventRepositoryDb() EventRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSslMode := os.Getenv("DB_SSLMODE")

	datasource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbAddress, dbPort, dbName, dbSslMode)
	opt, err := pg.ParseURL(datasource)

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
