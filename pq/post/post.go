package post

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() {
	var err error
	const connStr = "admin:@tcp(127.0.0.1:3306)/blogdb"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

type Post struct {
	ID       int
	Title    string
	Body     string
	Comments []Comment
}

func Insert(p *Post) error {
	r := db.QueryRow("INSERT INTO posts(title, body) VALUES (?,?)", p.Title, p.Body)
	err := r.Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil
}

func All() ([]Post, error) {
	var posts []Post
	rows, err := db.Query("SELECT id, title, body FROM posts order by id desc")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.Title, &p.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func FindByID(id int) (*Post, error) {
	row := db.QueryRow("SELECT id, title, body FROM posts WHERE id = ?", id)
	var p Post
	err := row.Scan(&p.ID, &p.Title, &p.Body)
	if err != nil {
		return nil, err
	}
	// Load Comments
	rows, err := db.Query("SELECT id, body, post_id FROM comments WHERE post_id = ? order by id desc", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c Comment
		err := rows.Scan(&c.ID, &c.Body, &c.PostID)
		if err != nil {
			return nil, err
		}
		p.Comments = append(p.Comments, c)
	}
	return &p, nil
}

func Save(p *Post) error {
	_, err := db.Exec("UPDATE posts SET title = ?, body = ? WHERE id = ?", p.Title, p.Body, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func AddComment(p *Post, c *Comment) error {
	r := db.QueryRow("INSERT INTO comments(body, post_id) VALUES (?,?)", c.Body, p.ID)
	err := r.Scan(&c.ID)
	if err != nil {
		return err
	}
	c.PostID = p.ID
	p.Comments = append(p.Comments, *c)
	return nil
}
