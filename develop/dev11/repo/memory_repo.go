package repo

import (
	"sync"
	"time"

	"github.com/Wmuga/wildberries-l2/develop/dev11/event"
)

type userEvents struct {
	lastID event.EventID
	events map[event.EventID]event.Event
}

type memoryRepo struct {
	mux    *sync.Mutex
	events map[event.UserID]*userEvents
}

type truncType int

const (
	truncDay truncType = iota
	truncWeek
	truncMonth
)

// AddEvent implements EventRepo.
func (m *memoryRepo) AddEvent(ev event.Event) (event.EventID, error) {
	m.mux.Lock()
	defer m.mux.Unlock()
	// Обновление ID добавление эвента
	userEvents := m.getOrCreateUser(ev.UserID)
	newID := event.GenerateNewID(userEvents.lastID)
	userEvents.lastID = newID
	ev.EventID = newID
	userEvents.events[newID] = ev

	return newID, nil
}

// GetEvent implements EventRepo.
func (m *memoryRepo) GetEvent(uid event.UserID, eid event.EventID) (ev event.Event, err error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	userEvents := m.getOrCreateUser(uid)
	ev, ex := userEvents.events[eid]
	if !ex {
		err = ErrEventNotFound
	}
	return
}

// UpdateEvent implements EventRepo.
func (m *memoryRepo) UpdateEvent(ev event.Event) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	userEvents := m.getOrCreateUser(ev.UserID)
	if _, ex := userEvents.events[ev.EventID]; !ex {
		return ErrEventNotFound
	}

	userEvents.events[ev.EventID] = ev
	return nil
}

// DeleteEvent implements EventRepo.
func (m *memoryRepo) DeleteEvent(userID event.UserID, evID event.EventID) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	userEvents := m.getOrCreateUser(userID)
	if _, ex := userEvents.events[evID]; !ex {
		return ErrEventNotFound
	}

	delete(userEvents.events, evID)
	return nil
}

// GetEventsDay implements EventRepo.
func (m *memoryRepo) GetEventsDay(userID event.UserID) ([]event.Event, error) {
	return m.getEventsPeriod(userID, truncDay, time.Now().Add(time.Hour*24)), nil
}

// GetEventsMonth implements EventRepo.
func (m *memoryRepo) GetEventsMonth(userID event.UserID) ([]event.Event, error) {
	return m.getEventsPeriod(userID, truncDay, time.Now().Add(time.Hour*24*31)), nil
}

// GetEventsWeek implements EventRepo.
func (m *memoryRepo) GetEventsWeek(userID event.UserID) ([]event.Event, error) {
	return m.getEventsPeriod(userID, truncDay, time.Now().Add(time.Hour*24*7)), nil
}

// getEventsPeriod Основная функция получения событий пользователя за данный период
func (m *memoryRepo) getEventsPeriod(userID event.UserID, ttype truncType, upTo time.Time) []event.Event {
	today := truncateDate(time.Now(), ttype)
	upTo = truncateDate(upTo, ttype)

	m.mux.Lock()
	defer m.mux.Unlock()
	userEvents := m.getOrCreateUser(userID)

	events := make([]event.Event, 0)
	for _, v := range userEvents.events {
		// Проверка, что евент начиная с сегодня,
		if !v.Date.Before(today) && v.Date.Before(upTo) {
			events = append(events, v)
		}
	}

	return events
}

// getOrCreateUser возвращает структуру пользовательских эвентов. Если ее нет - создает
func (m *memoryRepo) getOrCreateUser(id event.UserID) *userEvents {
	if ue, ex := m.events[id]; ex {
		return ue
	}
	ue := &userEvents{
		events: map[event.EventID]event.Event{},
	}
	m.events[id] = ue
	return ue
}

func NewMemoryRepo() EventRepo {
	return &memoryRepo{
		mux:    &sync.Mutex{},
		events: map[event.UserID]*userEvents{},
	}
}

// truncateDate квантует дату
func truncateDate(t time.Time, ttype truncType) time.Time {
	switch ttype {
	case truncDay:
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	case truncWeek:
		return time.Date(t.Year(), t.Month(), max((t.Day()/7), 1)*7, 0, 0, 0, 0, t.Location())
	case truncMonth:
		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	}
	return t
}
