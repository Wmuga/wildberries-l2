package service

import (
	"github.com/Wmuga/wildberries-l2/develop/dev11/event"
	"github.com/Wmuga/wildberries-l2/develop/dev11/repo"
)

// EventService сервис событий календаря
type EventService interface {
	AddEvent(event.Event) (event.EventID, error)
	UpdateEvent(event.Event) error
	DeleteEvent(event.UserID, event.EventID) error
	GetEventsDay(event.UserID) ([]event.Event, error)
	GetEventsWeek(event.UserID) ([]event.Event, error)
	GetEventsMonth(event.UserID) ([]event.Event, error)
}

type eventService struct {
	repo repo.EventRepo
}

// AddEvent implements EventService.
func (e *eventService) AddEvent(ev event.Event) (event.EventID, error) {
	return e.repo.AddEvent(ev)
}

// DeleteEvent implements EventService.
func (e *eventService) DeleteEvent(uid event.UserID, eid event.EventID) error {
	return e.repo.DeleteEvent(uid, eid)
}

// GetEventsDay implements EventService.
func (e *eventService) GetEventsDay(uid event.UserID) ([]event.Event, error) {
	return e.repo.GetEventsDay(uid)
}

// GetEventsMonth implements EventService.
func (e *eventService) GetEventsMonth(uid event.UserID) ([]event.Event, error) {
	return e.repo.GetEventsMonth(uid)
}

// GetEventsWeek implements EventService.
func (e *eventService) GetEventsWeek(uid event.UserID) ([]event.Event, error) {
	return e.repo.GetEventsWeek(uid)
}

// UpdateEvent implements EventService.
func (e *eventService) UpdateEvent(ev event.Event) error {
	return e.repo.UpdateEvent(ev)
}

func NewEventService(repo repo.EventRepo) EventService {
	return &eventService{repo}
}
