package service

import (
	"github.com/monicaribeiro/event-manager/domain"
	"github.com/monicaribeiro/event-manager/errs"
)

type EventService interface {
	GetAllEvents() ([]domain.Event, *errs.AppError)
	Delete(int642 int64) (*domain.Event, *errs.AppError)
}

type DefaultEventService struct {
	repo domain.EventRepository
}

func (s DefaultEventService) GetAllEvents() ([]domain.Event, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultEventService) Delete(id int64) (*domain.Event, *errs.AppError) {
	event, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func NewEventService(repository domain.EventRepository) DefaultEventService {
	return DefaultEventService{repository}
}
