package service

import "github.com/monicaribeiro/event-manager/domain"

type EventService interface {
	GetAllEvents() ([]domain.Event, error)
}

type DefaultEventService struct {
	repo domain.EventRepository
}

func (s DefaultEventService) GetAllEvents () ([]domain.Event, error) {
	return s.repo.FindAll()
}

func NewEventService(repository domain.EventRepository) DefaultEventService {
	return DefaultEventService{repository}
}