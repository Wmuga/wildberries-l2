package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Wmuga/wildberries-l2/develop/dev11/event"
	"github.com/Wmuga/wildberries-l2/develop/dev11/repo"
	"github.com/Wmuga/wildberries-l2/develop/dev11/service"
)

var (
	ev           = event.Event{EventID: 1, UserID: 1, Date: time.Now(), Description: "test"}
	url          = "http://localhost/test_url"
	urlWithQuery = "http://localhost/test_url?user_id=1"
)

func setupController() *EventsController {
	evRepo := repo.NewMemoryRepo()
	serv := service.NewEventService(evRepo)
	return NewEventController(serv, log.New(io.Discard, "", 0))
}

func testNoError(prefix string, r *http.Request, t *testing.T, f func(http.ResponseWriter, *http.Request)) {
	w := httptest.NewRecorder()
	f(w, r)
	data, err := io.ReadAll(w.Body)
	if err != nil {
		t.Error(prefix, err)
		return
	}

	resp := string(data)
	if strings.Contains(resp, "error") {
		t.Error(prefix, resp)
	}
}

func getBodyRequest(data []byte) *http.Request {
	return httptest.NewRequest("POST", url, bytes.NewBuffer(data))
}

func TestEventController(t *testing.T) {
	ctrl := setupController()
	data, err := ev.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}

	testNoError("Create", getBodyRequest(data), t, ctrl.CreateEvent)
	testNoError("Update", getBodyRequest(data), t, ctrl.UpdateEvent)
	testNoError("Delete", getBodyRequest(data), t, ctrl.DeleteEvent)
	r := httptest.NewRequest("GET", urlWithQuery, nil)
	testNoError("GetDay", r, t, ctrl.EventsForDay)
	testNoError("GetWeek", r, t, ctrl.EventsForWeek)
	testNoError("GetMonth", r, t, ctrl.EventsForMonth)
}
