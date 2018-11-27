package event

import "time"

type logAttendee struct {
	ID        int
	EventName string
	UserID    string
	Datetime  time.Time
}
