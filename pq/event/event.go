package event

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Event struct {
	ID             int
	Name           string
	Location       string
	Generation     int
	Speaker        string
	Description    string
	LimitAttendee  int
	AmountAttendee int
	StartDatetime  string
	EndDatetime    string
}

var db *sql.DB

func ConnectDB() {
	var err error
	const connStr = "admin:@tcp(127.0.0.1:3306)/kbtgevent?parseTime=true"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(e *Event) error {
	_, err := db.Exec("INSERT INTO event(name, location, generation, speaker, description, limit_attendee, amount_attendee, start_datetime, end_datetime) VALUES (?,?,?,?,?,?,?,?,?)", e.Name, e.Location, e.Generation, e.Speaker, e.Description, e.LimitAttendee, e.AmountAttendee, e.StartDatetime, e.EndDatetime)
	if err != nil {
		return err
	}
	return nil
}

func All() ([]Event, error) {
	var events []Event
	rows, err := db.Query("SELECT * FROM event order by id desc")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Location, &e.Generation, &e.Speaker, &e.Description, &e.LimitAttendee, &e.AmountAttendee, &e.StartDatetime, &e.EndDatetime)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func AllByEventName(eventName string) ([]Event, error) {
	var events []Event
	rows, err := db.Query("SELECT * FROM event WHERE name = ? order by id desc", eventName)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Location, &e.Generation, &e.Speaker, &e.Description, &e.LimitAttendee, &e.AmountAttendee, &e.StartDatetime, &e.EndDatetime)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func FindByID(id int) (*Event, error) {
	row := db.QueryRow("SELECT * FROM event WHERE id = ?", id)
	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Location, &e.Generation, &e.Speaker, &e.Description, &e.LimitAttendee, &e.AmountAttendee, &e.StartDatetime, &e.EndDatetime)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func Save(e *Event) error {
	_, err := db.Exec("UPDATE event SET name = ?, location = ?, generation = ?, speaker = ?, description = ?, limit_attendee = ?, amount_attendee = ?, start_datetime = ?, end_datetime = ? WHERE id = ?", e.Name, e.Location, e.Generation, e.Speaker, e.Description, e.LimitAttendee, e.AmountAttendee, e.StartDatetime, e.EndDatetime, e.ID)
	if err != nil {
		return err
	}
	return nil
}
