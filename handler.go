package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"web-event/pq/event"

	"github.com/Luxurioust/excelize"
)

var mu sync.Mutex

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	events, err := event.All()
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Events []event.Event
	}{
		Events: events,
	}
	err = indexTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func detailEventHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/events/"), "/detail"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	e, err := event.FindByID(id)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	events, err := event.AllByEventName(e.Name)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Events []event.Event
	}{
		Events: events,
	}
	err = detailEventTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func exportDataHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/events/"), "/export"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	e, err := event.FindByID(id)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	logAttendees, err := event.AllLog(e.Name, e.Generation)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		LogAttendees []event.LogAttendee
	}{
		LogAttendees: logAttendees,
	}
	// cols := []string{"A", "B", "C", "D", "E", "F"}
	xlsx := excelize.NewFile()
	// Set value of a cell.
	xlsx.SetCellValue("Sheet1", "A1", "ID")
	xlsx.SetCellValue("Sheet1", "B1", "Event name")
	xlsx.SetCellValue("Sheet1", "C1", "Generation")
	xlsx.SetCellValue("Sheet1", "D1", "User ID")
	xlsx.SetCellValue("Sheet1", "E1", "First name")
	xlsx.SetCellValue("Sheet1", "F1", "Last name")
	xlsx.SetCellValue("Sheet1", "G1", "Phone number")

	for index, atd := range data.LogAttendees {
		index := strconv.Itoa(index + 2)
		xlsx.SetCellValue("Sheet1", "A"+index, atd.ID)
		xlsx.SetCellValue("Sheet1", "B"+index, atd.EventName)
		xlsx.SetCellValue("Sheet1", "C"+index, atd.Generation)
		xlsx.SetCellValue("Sheet1", "D"+index, atd.UserID)
		xlsx.SetCellValue("Sheet1", "E"+index, atd.FirstName)
		xlsx.SetCellValue("Sheet1", "F"+index, atd.LastName)
		xlsx.SetCellValue("Sheet1", "G"+index, atd.PhoneNumber)
	}
	// Save xlsx file by the given path.
	genStr := strconv.Itoa(e.Generation)
	err = xlsx.SaveAs("./" + e.Name + "-" + genStr + "-attendees.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, fmt.Sprintf("/events/%d/edit", e.ID), http.StatusMovedPermanently)
}

func eventPageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/events/"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	e, err := event.FindByID(id)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = eventTemplate.Execute(w, e)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/events/"), "/update"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	e, err := event.FindByID(id)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
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
	e.Name = r.PostFormValue("name")
	e.Location = r.PostFormValue("location")
	e.Generation = generation
	e.Speaker = r.PostFormValue("speaker")
	e.Description = r.PostFormValue("description")
	e.LimitAttendee = limit
	e.StartDatetime = r.PostFormValue("start")
	e.EndDatetime = r.PostFormValue("end")
	err = event.Save(e)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/events/%d", e.ID), http.StatusMovedPermanently)
}

func joinEventPageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/events/"), "/join"))
	if err != nil {
		http.Error(w, "blog1: "+err.Error(), http.StatusInternalServerError)
		return
	}
	e, err := event.FindByID(id)
	if err != nil {
		http.Error(w, "blog2: "+err.Error(), http.StatusInternalServerError)
		return
	}
	e.AmountAttendee = e.AmountAttendee + 1
	newJoiner := &event.LogAttendee{
		EventName:   e.Name,
		Generation:  e.Generation,
		UserID:      r.PostFormValue("userid"),
		FirstName:   r.PostFormValue("firstname"),
		LastName:    r.PostFormValue("lastname"),
		PhoneNumber: r.PostFormValue("phonenumber"),
	}
	err = event.InsertLA(newJoiner)
	if err != nil {
		http.Error(w, "blog4: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = event.Save(e)
	if err != nil {
		http.Error(w, "blog3: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/events/%d", e.ID), http.StatusMovedPermanently)
}

func addEventHandler(w http.ResponseWriter, r *http.Request) {
	generation, err := strconv.Atoi(r.PostFormValue("generation"))
	if err != nil {
		http.Error(w, "blog gen: "+err.Error(), http.StatusInternalServerError)
		return
	}
	limit, err := strconv.Atoi(r.PostFormValue("limit"))
	if err != nil {
		http.Error(w, "blog limit: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newEvent := &event.Event{
		Name:          r.PostFormValue("name"),
		Location:      r.PostFormValue("location"),
		Generation:    generation,
		Speaker:       r.PostFormValue("speaker"),
		Description:   r.PostFormValue("description"),
		LimitAttendee: limit,
		StartDatetime: r.PostFormValue("start"),
		EndDatetime:   r.PostFormValue("end"),
	}
	err = event.Insert(newEvent)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/events/%d", newEvent.ID), http.StatusMovedPermanently)
}

func newEventPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := newEventTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func editEventPageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/events/"), "/edit"))
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	e, err := event.FindByID(id)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = editTemplate.Execute(w, e)
	if err != nil {
		http.Error(w, "blog: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func startServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/events/":
			indexPageHandler(w, r)
		case r.Method == http.MethodGet && r.URL.Path == "/events/new":
			newEventPageHandler(w, r)
		case r.Method == http.MethodPost && r.URL.Path == "/events/":
			addEventHandler(w, r)
		case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/events/") && strings.HasSuffix(r.URL.Path, "/join"):
			joinEventPageHandler(w, r)
		case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/events/") && strings.HasSuffix(r.URL.Path, "/edit"):
			editEventPageHandler(w, r)
		case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/events/") && strings.HasSuffix(r.URL.Path, "/update"):
			updateEventHandler(w, r)
		case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/events/") && strings.HasSuffix(r.URL.Path, "/export"):
			exportDataHandler(w, r)
		case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/events/") && strings.HasSuffix(r.URL.Path, "/detail"):
			detailEventHandler(w, r)
		case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/events/"):
			eventPageHandler(w, r)
		}
	})
	return http.ListenAndServe(":8000", nil)
}
