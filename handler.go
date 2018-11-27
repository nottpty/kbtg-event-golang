package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"web-event/pq/event"
	"web-event/pq/post"
)

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	posts, err := post.All()
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Posts []post.Post
	}{
		Posts: posts,
	}
	err = indexTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func postPageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/posts/"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	p, err := post.FindByID(id)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = postTemplate.Execute(w, p)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func addCommentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/posts/"), "/comment"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	p, err := post.FindByID(id)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var newComment post.Comment
	newComment.Body = r.PostFormValue("body")
	err = post.AddComment(p, &newComment)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/posts/%d", p.ID), http.StatusMovedPermanently)
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/posts/"), "/update"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	p, err := post.FindByID(id)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	p.Title = r.PostFormValue("title")
	p.Body = r.PostFormValue("body")
	err = post.Save(p)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/posts/%d", p.ID), http.StatusMovedPermanently)
}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	newPost := &post.Post{
		Title: r.PostFormValue("title"),
		Body:  r.PostFormValue("body"),
	}
	err := post.Insert(newPost)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/posts/%d", newPost.ID), http.StatusMovedPermanently)
}

func addEventHandler(w http.ResponseWriter, r *http.Request) {
	generation, err := strconv.Atoi(r.PostFormValue("generation"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	limit, err := strconv.Atoi(r.PostFormValue("limit"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	startDate := time.Now()
	endDate := time.Now()
	newEvent := &event.Event{
		Name:          r.PostFormValue("name"),
		Location:      r.PostFormValue("location"),
		Generation:    generation,
		Speaker:       r.PostFormValue("speaker"),
		Description:   r.PostFormValue("description"),
		LimitAttendee: limit,
		StartDatetime: startDate,
		EndDatetime:   endDate,
	}
	err = event.Insert(newEvent)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/posts/%d", newEvent.ID), http.StatusMovedPermanently)
}

func newPostPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := newTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func newEventPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := newEventTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func editPostPageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/posts/"), "/edit"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	p, err := post.FindByID(id)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = editTemplate.Execute(w, p)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func startServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/posts/":
			indexPageHandler(w, r)
		case r.Method == http.MethodGet && r.URL.Path == "/posts/new":
			newPostPageHandler(w, r)
		case r.Method == http.MethodGet && r.URL.Path == "/events/new":
			newEventPageHandler(w, r)
		case r.Method == http.MethodPost && r.URL.Path == "/posts/":
			addPostHandler(w, r)
		case r.Method == http.MethodPost && r.URL.Path == "/events/":
			addEventHandler(w, r)
		case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/posts/") && strings.HasSuffix(r.URL.Path, "/comment"):
			addCommentHandler(w, r)
		case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/posts/") && strings.HasSuffix(r.URL.Path, "/edit"):
			editPostPageHandler(w, r)
		case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/posts/") && strings.HasSuffix(r.URL.Path, "/update"):
			updatePostHandler(w, r)
		case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/posts/"):
			postPageHandler(w, r)
		}
	})
	return http.ListenAndServe(":8000", nil)
}
