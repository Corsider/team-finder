package usecase

import (
	"team-finder/domain"
	"time"
)

type eventUsecase struct {
	eventRepository domain.EventRepository
	contextTimeout  time.Duration
}

func NewEventUsecase(eventRepository domain.EventRepository, timeout time.Duration) domain.EventUsecase {
	return &eventUsecase{
		eventRepository: eventRepository,
		contextTimeout:  timeout,
	}
}

func (eu *eventUsecase) GetAll() ([]domain.Event, error) {
	return eu.eventRepository.GetAll()
}

func (eu *eventUsecase) GetEventById(eventId int) (domain.Event, error) {
	return eu.eventRepository.GetEventById(eventId)
}

func (eu *eventUsecase) RegEvent(request domain.EventRegRequest, creatorId int) (int, error) {
	return eu.eventRepository.RegEvent(request, creatorId)
}
