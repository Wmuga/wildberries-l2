package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Wmuga/wildberries-l2/develop/dev11/event"
	"github.com/Wmuga/wildberries-l2/develop/dev11/service"
)

type EventsController struct {
	service service.EventService
	logger  *log.Logger
}

type EventGetter func(event.UserID) ([]event.Event, error)

var (
	strSuccess = "success"
)

func NewEventController(service service.EventService, errLogger *log.Logger) *EventsController {
	return &EventsController{
		service,
		errLogger,
	}
}

func (e *EventsController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	// Чтение POST
	ev, ok := e.parseEvent(w, r)
	if !ok {
		return
	}
	// Добавить в сервис
	id, err := e.service.AddEvent(ev)
	if err != nil {
		e.errorResponse(err, http.StatusServiceUnavailable, w)
		return
	}
	e.resultResponse(event.EventIDToString(id), w)
}

func (e *EventsController) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	// Чтение POST
	ev, ok := e.parseEvent(w, r)
	if !ok {
		return
	}
	// Добавить в сервис
	err := e.service.UpdateEvent(ev)
	if err != nil {
		e.errorResponse(err, http.StatusServiceUnavailable, w)
		return
	}
	e.resultResponse(strSuccess, w)
}

func (e *EventsController) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	// Чтение POST
	ev, ok := e.parseEvent(w, r)
	if !ok {
		return
	}
	// Добавить в сервис
	err := e.service.DeleteEvent(ev.UserID, ev.EventID)
	if err != nil {
		e.errorResponse(err, http.StatusServiceUnavailable, w)
		return
	}
	e.resultResponse(strSuccess, w)
}

func (e *EventsController) EventsForDay(w http.ResponseWriter, r *http.Request) {
	e.sendEvents(w, r, e.service.GetEventsDay)
}

func (e *EventsController) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	e.sendEvents(w, r, e.service.GetEventsMonth)
}

func (e *EventsController) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	e.sendEvents(w, r, e.service.GetEventsMonth)
}

func (e *EventsController) sendEvents(w http.ResponseWriter, r *http.Request, getter EventGetter) {
	// Получение uid
	uidStr := r.URL.Query().Get("user_id")
	uid, err := event.StringToUserID(uidStr)
	if err != nil {
		e.errorResponse(err, http.StatusBadRequest, w)
		return
	}
	// Получение данных
	events, err := getter(uid)
	if err != nil {
		e.errorResponse(err, http.StatusServiceUnavailable, w)
		return
	}
	data, err := event.MarshalJSON(events)
	if err != nil {
		e.errorResponse(err, http.StatusServiceUnavailable, w)
		return
	}
	e.resultResponse(string(data), w)
}

func readBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

func (e *EventsController) parseEvent(w http.ResponseWriter, r *http.Request) (ev event.Event, ok bool) {
	body, err := readBody(r)
	if err != nil {
		e.errorResponse(err, http.StatusInternalServerError, w)
		return
	}
	ev, err = event.UnmarshalJSON(body)
	if err != nil {
		e.errorResponse(err, http.StatusBadRequest, w)
		return
	}
	ok = true
	return
}

func (e *EventsController) errorResponse(err error, statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	_, err = w.Write([]byte(fmt.Sprintf(`"error":"%s"`, err)))
	if err != nil {
		e.logger.Println(err)
	}
}

func (e *EventsController) resultResponse(result string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf(`"result":"%s"`, result)))
	if err != nil {
		e.logger.Println(err)
	}
}
