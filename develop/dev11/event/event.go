package event

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type UserID uint
type EventID uint

type Event struct {
	EventID     EventID
	UserID      UserID
	Date        time.Time
	Description string
}

// eventPost структура для первоначального парса POST данных
type eventPost struct {
	EventID     EventID `json:"event_id"`
	UserID      UserID  `json:"user_id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
}

var (
	ErrUnknownDate = errors.New("wrong date format")
)

// GenerateNewID создает новый id на основе предыдущего
func GenerateNewID(id EventID) EventID {
	return id + 1
}

// Event.MarshalJSON представляет Event в виде строки формата JSON
func (e Event) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}

	if _, err := buf.WriteRune('{'); err != nil {
		return nil, err
	}

	if _, err := buf.WriteString(
		fmt.Sprintf(`"event_id":%d,`, e.EventID)); err != nil {
		return nil, err
	}

	if _, err := buf.WriteString(
		fmt.Sprintf(`"user_id":%d,`, e.UserID),
	); err != nil {
		return nil, err
	}

	if _, err := buf.WriteString(
		fmt.Sprintf(`"description":"%s",`, e.Description),
	); err != nil {
		return nil, err
	}

	if _, err := buf.WriteString(fmt.Sprintf(`"date":"%d-%02d-%02d"`,
		e.Date.Year(), e.Date.Month(), e.Date.Day()),
	); err != nil {
		return nil, err
	}

	if _, err := buf.WriteRune('}'); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSON предаставляет массив Event'ов в виде JSON
func MarshalJSON(evs []Event) ([]byte, error) {
	if len(evs) == 0 {
		return []byte("[]"), nil
	}

	buf := bytes.Buffer{}

	if _, err := buf.WriteRune('['); err != nil {
		return nil, err
	}

	// отдельно первое событие
	data, err := evs[0].MarshalJSON()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(data); err != nil {
		return nil, err
	}

	// остальные через запятую
	for i := 1; i < len(evs); i++ {
		if _, err := buf.WriteRune(','); err != nil {
			return nil, err
		}
		data, err := evs[i].MarshalJSON()
		if err != nil {
			return nil, err
		}
		if _, err := buf.Write(data); err != nil {
			return nil, err
		}
	}

	if _, err := buf.WriteRune(']'); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON предобразует JSON строку в объект Event
func UnmarshalJSON(data []byte) (e Event, err error) {
	postEvent := &eventPost{}
	err = json.Unmarshal(data, postEvent)
	if err != nil {
		return
	}
	// копирование полей
	e.EventID = postEvent.EventID
	e.UserID = postEvent.UserID
	e.Description = postEvent.Description
	// скип парса даты, если она пустая
	if postEvent.Date == "" {
		return
	}
	// парс строки даты
	dateSplit := strings.Split(postEvent.Date, "-")
	if len(dateSplit) < 3 {
		err = ErrUnknownDate
		return
	}
	var dateInt [3]int
	for i := range dateSplit {
		dateInt[i], err = strconv.Atoi(dateSplit[i])
		if err != nil {
			return
		}
	}
	e.Date = time.Date(dateInt[0], time.Month(dateInt[1]), dateInt[2], 0, 0, 0, 0, time.Local)

	return
}

func EventIDToString(id EventID) string {
	return strconv.Itoa(int(id))
}

func StringToUserID(uidStr string) (UserID, error) {
	id, err := strconv.Atoi(uidStr)
	return UserID(id), err
}
