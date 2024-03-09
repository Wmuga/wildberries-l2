package repo

import (
	"reflect"
	"testing"
	"time"

	"github.com/Wmuga/wildberries-l2/develop/dev11/event"
)

var (
	today    = time.Now()
	tomorrow = time.Date(today.Year(), today.Month(), today.Day()+1, 0, 0, 0, 0, today.Location())
	nextWeek = time.Date(today.Year(), today.Month(), today.Day()+7, 0, 0, 0, 0, today.Location())
	events   = []event.Event{
		{EventID: 0, UserID: 1, Date: today, Description: "Test1"},
		{EventID: 0, UserID: 1, Date: tomorrow, Description: "Test2"},
		{EventID: 0, UserID: 1, Date: nextWeek, Description: "Test3"},
		{EventID: 0, UserID: 2, Date: today, Description: "Test4"}}
)

func setupMemoryEvents() (repo EventRepo, err error) {
	repo = NewMemoryRepo()
	for i := range events {
		var id event.EventID
		id, err = repo.AddEvent(events[i])
		if err != nil {
			return
		}
		events[i].EventID = id
	}
	return
}

func testGetterBase(expSize int, t *testing.T, getter func(event.UserID) ([]event.Event, error)) {
	exp := events[:expSize]
	got, err := getter(events[0].UserID)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("wrong events\n\t Got: %#v\n\tWant: %#v", got, exp)
	}
}

func testRepoGetUpdate(repo EventRepo, t *testing.T) {
	// Проверка получения отдельных событий
	for _, ev := range events {
		// Проверка получения
		ev2, err := repo.GetEvent(ev.UserID, ev.EventID)
		if err != nil {
			t.Error(err)
			continue
		}

		if ev != ev2 {
			t.Errorf("wrong event\n\t Got: %#v\n\tWant: %#v", ev2, ev)
		}
	}
	// Проверка обновления
	ev := events[3]
	ev.Description = "new description"
	err := repo.UpdateEvent(ev)
	if err != nil {
		t.Error(err)
	}
	ev2, err := repo.GetEvent(ev.UserID, ev.EventID)
	if err != nil {
		t.Error(err)
	}
	if ev != ev2 {
		t.Errorf("didn't update event\n\t Got: %#v\n\tWant: %#v", ev2, ev)
	}
}

func testRepoGetters(repo EventRepo, t *testing.T) {
	testGetterBase(1, t, repo.GetEventsDay)
	testGetterBase(2, t, repo.GetEventsWeek)
	testGetterBase(3, t, repo.GetEventsMonth)
}

func TestMemoryRepo(t *testing.T) {
	// Setup с добавлением эвентов
	repo, err := setupMemoryEvents()
	if err != nil {
		t.Error(err)
		return
	}
	testRepoGetUpdate(repo, t)
	testRepoGetters(repo, t)
}
