package event

import (
	"database/sql"
	"log"
	"time"

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
	StartDatetime  time.Time
	EndDatetime    time.Time
}

var db *sql.DB

func ConnectDB() {
	var err error
	const connStr = "admin:@tcp(127.0.0.1:3306)/kbtgevent"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(e *Event) error {
	r := db.QueryRow("INSERT INTO event(name, location, generation, speaker, description, limit_attendee, amount_attendee, start_datetime, end_datetime) VALUES (?,?,?,?,?,?,?,?,?)", e.Name, e.Location, e.Generation, e.Speaker, e.Description, e.LimitAttendee, e.AmountAttendee, e.StartDatetime, e.EndDatetime)
	err := r.Scan(&e.ID)
	if err != nil {
		return err
	}
	return nil
}

// func All() ([]Event, error) {
// 	var posts []Event
// 	rows, err := db.Query("SELECT id, title, body FROM posts order by id desc")
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var p Event
// 		err := rows.Scan(&p.ID, &p.Title, &p.Body)
// 		if err != nil {
// 			return nil, err
// 		}
// 		posts = append(posts, p)
// 	}
// 	return posts, nil
// }

// func FindByID(id int) (*Event, error) {
// 	row := db.QueryRow("SELECT id, title, body FROM posts WHERE id = ?", id)
// 	var p Event
// 	err := row.Scan(&p.ID, &p.Title, &p.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Load Comments
// 	rows, err := db.Query("SELECT id, body, post_id FROM comments WHERE post_id = ? order by id desc", id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var c Comment
// 		err := rows.Scan(&c.ID, &c.Body, &c.PostID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		p.Comments = append(p.Comments, c)
// 	}
// 	return &p, nil
// }

// func Save(p *Event) error {
// 	_, err := db.Exec("UPDATE posts SET title = ?, body = ? WHERE id = ?", p.Title, p.Body, p.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func AddComment(p *Event, c *Comment) error {
// 	r := db.QueryRow("INSERT INTO comments(body, post_id) VALUES (?,?)", c.Body, p.ID)
// 	err := r.Scan(&c.ID)
// 	if err != nil {
// 		return err
// 	}
// 	c.PostID = p.ID
// 	p.Comments = append(p.Comments, *c)
// 	return nil
// }
