package event

type LogAttendee struct {
	ID          int
	EventName   string
	UserID      string
	FirstName   string
	LastName    string
	PhoneNumber string
}

func InsertLA(la *LogAttendee) error {
	_, err := db.Exec("INSERT INTO log_attendee(event_name, user_id, first_name, last_name, phone_number) VALUES (?,?,?,?,?)", la.EventName, la.UserID, la.FirstName, la.LastName, la.PhoneNumber)
	// err = r.Scan(&la.ID)
	if err != nil {
		return err
	}
	return nil
}

func AllLog(eventName string) ([]LogAttendee, error) {
	var logAttendees []LogAttendee
	rows, err := db.Query("SELECT * FROM log_attendee WHERE event_name = ? order by id asc", eventName)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var la LogAttendee
		err := rows.Scan(&la.ID, &la.EventName, &la.UserID, &la.FirstName, &la.LastName, &la.PhoneNumber)
		if err != nil {
			return nil, err
		}
		logAttendees = append(logAttendees, la)
	}
	return logAttendees, nil
}

// type user struct {
// 	UserID      string
// 	FirsName    string
// 	LastName    string
// 	PhoneNumber string
// }
