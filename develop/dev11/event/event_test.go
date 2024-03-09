package event

import (
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	date := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local)
	expected := `{"event_id":1,"user_id":2,"description":"test","date":"2023-01-01"}`
	expectedArr := `[{"event_id":1,"user_id":2,"description":"test","date":"2023-01-01"},{"event_id":1,"user_id":2,"description":"test","date":"2023-01-01"}]`

	ev := Event{
		1, 2, date, "test",
	}
	evs := []Event{ev, ev}

	data, err := ev.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}

	if string(data) != expected {
		t.Errorf("Error in unmarshalling\n\tGot %#v\n\tExpected %#v", string(data), expected)
		return
	}

	ev2, err := UnmarshalJSON(data)
	if err != nil {
		t.Error(err)
		return
	}

	if ev != ev2 {
		t.Errorf("Error in unmarshalling\n\tGot %#v\n\tExpected %#v", ev2, ev)
	}

	data, err = MarshalJSON(evs)
	if err != nil {
		t.Error(err)
		return
	}

	if string(data) != expectedArr {
		t.Errorf("Error in unmarshalling\n\tGot %#v\n\tExpected %#v", string(data), expectedArr)
		return
	}
}
