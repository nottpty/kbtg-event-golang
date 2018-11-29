package event

type LogAttendee struct {
	ID          int
	EventName   string
	Generation  int
	UserID      string
	FirstName   string
	LastName    string
	PhoneNumber string
}

func InsertLA(la *LogAttendee) error {
	_, err := db.Exec("INSERT INTO log_attendee(event_name, generation, user_id, first_name, last_name, phone_number) VALUES (?,?,?,?,?)", la.EventName, la.Generation, la.UserID, la.FirstName, la.LastName, la.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func AllLog(eventName string, gen int) ([]LogAttendee, error) {
	var logAttendees []LogAttendee
	rows, err := db.Query("SELECT * FROM log_attendee WHERE event_name = ? and generation = ? order by id asc", eventName, gen)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var la LogAttendee
		err := rows.Scan(&la.ID, &la.EventName, &la.Generation, &la.UserID, &la.FirstName, &la.LastName, &la.PhoneNumber)
		if err != nil {
			return nil, err
		}
		logAttendees = append(logAttendees, la)
	}
	return logAttendees, nil
}
