package repo

import (
	"errors"

	"github.com/Wmuga/wildberries-l2/develop/dev11/event"
)

var (
	ErrEventNotFound = errors.New("event not found")
)

type EventRepo interface {
	AddEvent(event.Event) (event.EventID, error)
	GetEvent(event.UserID, event.EventID) (event.Event, error)
	UpdateEvent(event.Event) error
	DeleteEvent(event.UserID, event.EventID) error
	GetEventsDay(event.UserID) ([]event.Event, error)
	GetEventsWeek(event.UserID) ([]event.Event, error)
	GetEventsMonth(event.UserID) ([]event.Event, error)
}
