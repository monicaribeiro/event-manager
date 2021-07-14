package service

import (
	"github.com/monicaribeiro/event-manager/domain"
	"github.com/monicaribeiro/event-manager/dto"
	"github.com/monicaribeiro/event-manager/errs"
	"github.com/monicaribeiro/event-manager/logger"
)

type EventService interface {
	GetAllEvents(state string) ([]dto.EventResponse, *errs.AppError)
	Delete(int642 int64) (*dto.EventResponse, *errs.AppError)
	Create(request dto.NewEventRequest) *errs.AppError
}

type DefaultEventService struct {
	repo domain.EventRepository
}

func (s DefaultEventService) GetAllEvents(state string) ([]dto.EventResponse, *errs.AppError) {
	logger.Info("V2 Getting all events")
	events, err := s.repo.FindAll(state)

	if err != nil {
		return nil, err
	}

	response := make([]dto.EventResponse, 0)
	for _, e := range events {
		response = append(response, e.ToDto())
	}
	return response, err
}

func (s DefaultEventService) Delete(id int64) (*dto.EventResponse, *errs.AppError) {
	logger.Info("V2 Deleting event")
	event, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(id)

	if err != nil {
		return nil, err
	}

	eventDto := event.ToDto()
	return &eventDto, nil
}

func (s DefaultEventService) Create(request dto.NewEventRequest) *errs.AppError {
	logger.Info("V2 Creating event")
	event := domain.NewEvent(request.Name, request.City, request.State, request.PhotoUrl, request.Datetime)
	err := s.repo.Create(&event)

	if err != nil {
		return err
	}

	return nil
}

func NewEventService(repository domain.EventRepository) DefaultEventService {
	return DefaultEventService{repository}
}
