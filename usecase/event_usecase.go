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

func (eu *eventUsecase) AddTeamToEvent(eventId, teamId int) error {
	return eu.eventRepository.AddTeamToEvent(eventId, teamId)
}

func (eu *eventUsecase) DeleteEvent(eventId int) error {

	err := eu.eventRepository.DeleteFromEventsTags(eventId)
	if err != nil {
		return err
	}
	err = eu.eventRepository.DeleteFromTeamEvent(eventId)
	if err != nil {
		return err
	}
	err = eu.eventRepository.DeleteFromEvents(eventId)
	if err != nil {
		return err
	}
	return nil
}
