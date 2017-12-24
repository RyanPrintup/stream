package stream

import "time"

type Event struct {
	Type string
	Timestamp string
	Args ...interface{}
}

func NewEvent(type string, args ...interface{}) *Event {
	return &Event{type, time.Now().UTC().Format(time.RFC3339), args};
}